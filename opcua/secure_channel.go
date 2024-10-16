package opcua

import (
	"errors"
	"sync/atomic"

	"golang.org/x/exp/slog"

	"github.com/protocol-laboratory/opcua-go/opcua/enc"
	"github.com/protocol-laboratory/opcua-go/opcua/uamsg"
	"github.com/protocol-laboratory/opcua-go/opcua/util"
)

type SecureChannel struct {
	conn         *opcuaConn
	channelId    uint32
	channelIdGen *ChannelIdGen
	logger       *slog.Logger

	receiveMaxChunkSize    uint32
	sendMaxChunkSize       uint32
	maxChunkCount          uint32
	maxResponseMessageSize uint32
	endpointUrl            string

	// TODO conn set read timeout
	decoder enc.Decoder
	encoder enc.Encoder

	currentTokenId     atomic.Uint32
	nextSequenceNumber atomic.Uint32
}

const (
	// ProtocolVersion https://reference.opcfoundation.org/Core/Part6/v105/docs/7.1.2.3
	ProtocolVersion uint32 = 0
)

func newSecureChannel(conn *opcuaConn, svcConf *ServerConfig, channelId uint32, channelIdGen *ChannelIdGen, logger *slog.Logger) *SecureChannel {
	return &SecureChannel{
		conn:         conn,
		channelId:    channelId,
		channelIdGen: channelIdGen,
		logger:       logger,
		decoder:      enc.NewDefaultDecoder(conn.conn, int64(svcConf.ReceiverBufferSize)),
		encoder:      enc.NewDefaultEncoder(),
	}
}

func (secChan *SecureChannel) open() error {
	secChan.logger.Info("handling hello message")
	err := secChan.handleHello()
	if err != nil {
		return err
	}

	secChan.logger.Info("handling open secure channel message")
	err = secChan.handleOpenSecureChannel()
	if err != nil {
		return err
	}
	return nil
}

func (secChan *SecureChannel) handleHello() error {
	req, err := secChan.decoder.ReadMsg()
	if err != nil {
		return err
	}

	if req.MessageHeader.MessageType != uamsg.HelloMessageType {
		return errors.New("invalid message type, expected Hello")
	}

	helloBody, ok := req.MessageBody.(*uamsg.HelloMessageExtras)
	if !ok {
		return errors.New("invalid message body, expected HelloMessageExtras")
	}

	// TODO Instantiation new encoder and decoder
	secChan.receiveMaxChunkSize = helloBody.SendBufferSize
	secChan.sendMaxChunkSize = helloBody.ReceiveBufferSize
	secChan.maxChunkCount = helloBody.MaxChunkCount
	secChan.maxResponseMessageSize = helloBody.MaxMessageSize
	// TODO need callback to validate endpoint
	secChan.endpointUrl = helloBody.EndpointUrl

	resp := &uamsg.Message{
		MessageHeader: &uamsg.MessageHeader{
			MessageType: uamsg.AcknowledgeMessageType,
		},
		MessageBody: &uamsg.AcknowledgeMessageExtras{
			ProtocolVersion:   ProtocolVersion,
			ReceiveBufferSize: secChan.receiveMaxChunkSize,
			SendBufferSize:    secChan.sendMaxChunkSize,
			MaxMessageSize:    secChan.maxResponseMessageSize,
			MaxChunkCount:     secChan.maxChunkCount,
		},
	}

	bytes, err := secChan.encoder.Encode(resp, int(secChan.maxResponseMessageSize))
	if err != nil {
		return err
	}

	for _, content := range bytes {
		length, err := secChan.conn.conn.Write(content)
		if err != nil {
			return err
		}
		if length != len(content) {
			return errors.New("write length not match")
		}
	}

	return nil
}

func (secChan *SecureChannel) handleOpenSecureChannel() error {
	// TODO should return ERROR STATUS CODE to client if any error occur
	req, err := secChan.decoder.ReadMsg()
	if err != nil {
		return err
	}

	if req.MessageHeader.MessageType != uamsg.OpenSecureChannelMessageType {
		return errors.New("invalid message type, expected OpenSecureChannel")
	}

	genericBody, ok := req.MessageBody.(*uamsg.GenericBody)
	if !ok {
		return errors.New("invalid message body, expected GenericBody")
	}

	openSecChanBody, ok := genericBody.Service.(*uamsg.OpenSecureChannelServiceRequest)
	if !ok {
		return errors.New("invalid service, expected OpenSecureChannelServiceRequest")
	}

	// TODO only support NONE mode yet
	if openSecChanBody.SecurityMode != uamsg.MessageSecurityModeNone {
		return errors.New("only support NONE security mode")
	}

	serverNonce := []byte{}
	channelId := secChan.channelIdGen.next()
	tokenId := secChan.getNextTokenId()

	rsp := &uamsg.Message{
		MessageHeader: &uamsg.MessageHeader{
			MessageType:     uamsg.OpenSecureChannelMessageType,
			SecureChannelId: &channelId,
		},
		SecurityHeader: &uamsg.AsymmetricSecurityHeader{
			SecurityPolicyUri:             []byte("http://opcfoundation.org/UA/SecurityPolicy#None"),
			SenderCertificate:             nil,
			ReceiverCertificateThumbprint: nil,
		},
		SequenceHeader: &uamsg.SequenceHeader{
			SequenceNumber: secChan.getNextSequenceNumber(),
			RequestId:      req.RequestId,
		},
		MessageBody: &uamsg.GenericBody{
			TypeId: &uamsg.ExpandedNodeId{
				NodeId: &uamsg.ObjectOpenSecureChannelResponse_Encoding_DefaultBinary,
			},
			Service: &uamsg.OpenSecureChannelServiceResponse{
				Header: &uamsg.ResponseHeader{
					Timestamp:     util.GetCurrentUaTimestamp(),
					RequestHandle: openSecChanBody.Header.RequestHandle,
					ServiceResult: 0,
					ServiceDiagnostics: &uamsg.DiagnosticInfo{
						EncodingMask: 0x00,
					},
					StringTable: nil,
					AdditionalHeader: &uamsg.ExtensionObject{
						TypeId: &uamsg.NodeId{
							EncodingType: uamsg.TwoByte,
							Identifier:   byte(0),
						},
						Encoding: 0x00,
					},
				},
				ServerProtocolVersion: 0,
				SecurityToken: &uamsg.ChannelSecurityToken{
					ChannelID:       channelId,
					TokenID:         tokenId,
					CreatedAt:       util.GetCurrentUaTimestamp(),
					RevisedLifetime: 3600000,
				},
				ServerNonce: serverNonce,
			},
		},
	}

	return secChan.sendResponse(rsp)
}

func (secChan *SecureChannel) serve() error {
	for {
		req, err := secChan.decoder.ReadMsg()
		if err != nil {
			return err
		}

		err = secChan.handleRequest(req)
		if err != nil {
			return err
		}
	}
}

func (secChan *SecureChannel) handleRequest(req *uamsg.Message) error {
	generic, ok := req.MessageBody.(*uamsg.GenericBody)
	if !ok {
		return errors.New("invalid message body")
	}

	var rsp *uamsg.Message
	var err error
	switch generic.Service.(type) {
	case *uamsg.OpenSecureChannelServiceRequest:
		rsp, err = secChan.handleOpenSecureChannelServiceRequest(req)
	case *uamsg.CloseSecureChannelRequest:
		rsp, err = secChan.handleCloseSecureChannelRequest(req)
	case *uamsg.ReadRequest:
		rsp, err = secChan.handleReadRequest(req)
	case *uamsg.CreateSessionRequest:
		rsp, err = secChan.handleCreateSessionRequest(req)
	case *uamsg.ActivateSessionRequest:
		rsp, err = secChan.handleActivateSessionRequest(req)
	case *uamsg.CloseSessionRequest:
		rsp, err = secChan.handleCloseSessionRequest(req)
	default:
		secChan.logger.Error("unsupported an service", "service", generic.Service)
		rsp = secChan.createErrorMessage(req.RequestId, 0, uamsg.ErrorCodeBadServiceUnsupported)
	}

	if err != nil {
		return err
	}
	return secChan.sendResponse(rsp)
}

func (secChan *SecureChannel) sendResponse(rsp *uamsg.Message) error {
	bytes, err := secChan.encoder.Encode(rsp, int(secChan.maxResponseMessageSize))
	if err != nil {
		return err
	}

	for _, content := range bytes {
		length, err := secChan.conn.conn.Write(content)
		if err != nil {
			return err
		}
		if length != len(content) {
			return errors.New("write length not match")
		}
	}
	return nil
}

func (secChan *SecureChannel) createErrorMessage(reqId uint32, reqHandle uamsg.IntegerId, ec uamsg.ErrorCode) *uamsg.Message {
	return &uamsg.Message{
		MessageHeader: &uamsg.MessageHeader{
			MessageType:     uamsg.MsgMessageType,
			SecureChannelId: &secChan.channelId,
		},
		SecurityHeader: &uamsg.SymmetricSecurityHeader{
			TokenId: secChan.getCurrentTokenId(),
		},
		SequenceHeader: &uamsg.SequenceHeader{
			SequenceNumber: secChan.getNextSequenceNumber(),
			RequestId:      reqId,
		},
		MessageBody: &uamsg.GenericBody{
			TypeId: &uamsg.ExpandedNodeId{
				NodeId: &uamsg.ObjectServiceFault_Encoding_DefaultBinary,
			},
			Service: &uamsg.ServiceFault{
				Header: &uamsg.ResponseHeader{
					Timestamp:     util.GetCurrentUaTimestamp(),
					RequestHandle: reqHandle,
					ServiceResult: uamsg.StatusCode(ec),
					ServiceDiagnostics: &uamsg.DiagnosticInfo{
						EncodingMask: 0x00,
					},
					StringTable: nil,
					AdditionalHeader: &uamsg.ExtensionObject{
						TypeId: &uamsg.NodeId{
							EncodingType: uamsg.TwoByte,
							Identifier:   byte(0),
						},
						Encoding: 0x00,
					},
				},
			},
		},
	}
}

func (secChan *SecureChannel) getCurrentTokenId() uint32 {
	return secChan.currentTokenId.Load()
}

func (secChan *SecureChannel) getNextTokenId() uint32 {
	return secChan.currentTokenId.Add(1)
}

func (secChan *SecureChannel) getNextSequenceNumber() uint32 {
	return secChan.nextSequenceNumber.Add(1)
}
