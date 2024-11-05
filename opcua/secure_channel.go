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
	conn           *Conn
	channelId      uint32
	sessionManager *SessionManager
	interceptor    ServerInterceptor
	handler        ServerHandler
	logger         *slog.Logger

	readRequestNodeLimit uint32

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

func newSecureChannel(conn *Conn,
	svcConf *ServerConfig,
	channelId uint32,
	sessionManager *SessionManager,
	interceptor ServerInterceptor,
	handler ServerHandler,
	logger *slog.Logger) *SecureChannel {
	return &SecureChannel{
		conn:                 conn,
		channelId:            channelId,
		sessionManager:       sessionManager,
		interceptor:          interceptor,
		handler:              handler,
		readRequestNodeLimit: uint32(svcConf.ReadRequestNodeLimit),
		logger:               logger,
		decoder:              enc.NewDefaultDecoder(conn, int64(svcConf.ReceiverBufferSize)),
		encoder:              enc.NewDefaultEncoder(svcConf.MaxResponseSize),
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
	secChan.endpointUrl = helloBody.EndpointUrl

	err = secChan.interceptor.BeforeHello(secChan.conn, helloBody)
	if err != nil {
		return err
	}

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

	err = secChan.interceptor.AfterHello(secChan.conn, helloBody)
	if err != nil {
		return err
	}

	for _, content := range bytes {
		length, err := secChan.conn.Write(content)
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

	svc, ok := genericBody.Service.(*uamsg.OpenSecureChannelRequest)
	if !ok {
		return errors.New("invalid service, expected OpenSecureChannelRequest")
	}

	// TODO only support NONE mode yet
	if svc.SecurityMode != uamsg.MessageSecurityModeNone {
		return errors.New("only support NONE security mode")
	}

	var serverNonce []byte
	tokenId := secChan.getNextTokenId()

	respSvc := &uamsg.OpenSecureChannelResponse{
		Header: &uamsg.ResponseHeader{
			Timestamp:     util.GetCurrentUaTimestamp(),
			RequestHandle: svc.Header.RequestHandle,
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
			ChannelID:       secChan.channelId,
			TokenID:         tokenId,
			CreatedAt:       util.GetCurrentUaTimestamp(),
			RevisedLifetime: 3600000,
		},
		ServerNonce: serverNonce,
	}

	err = secChan.interceptor.AfterOpenSecureChannel(secChan.conn, svc, respSvc)
	if err != nil {
		return err
	}

	rsp := &uamsg.Message{
		MessageHeader: &uamsg.MessageHeader{
			MessageType:     uamsg.OpenSecureChannelMessageType,
			SecureChannelId: &secChan.channelId,
		},
		SecurityHeader: &uamsg.AsymmetricSecurityHeader{
			SecurityPolicyUri:             []byte(uamsg.SecurityPolicyUriNone),
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
			Service: respSvc,
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

	// TODO check channel id
	// TODO check token id
	// TODO check session (should discard requests on the closed session)

	var rsp *uamsg.Message
	var err error
	switch generic.Service.(type) {
	case *uamsg.OpenSecureChannelRequest:
		rsp, err = secChan.handleOpenSecureChannelRequest(req)
	case *uamsg.CloseSecureChannelRequest:
		rsp, err = secChan.handleCloseSecureChannelRequest(req)
	case *uamsg.CreateSessionRequest:
		rsp, err = secChan.handleCreateSessionRequest(req)
	case *uamsg.ActivateSessionRequest:
		rsp, err = secChan.handleActivateSessionRequest(req)
	case *uamsg.CloseSessionRequest:
		rsp, err = secChan.handleCloseSessionRequest(req)
	case *uamsg.GetEndpointsRequest:
		rsp, err = secChan.handleGetEndpoints(req)
	case *uamsg.BrowseRequest:
		rsp, err = secChan.handleBrowseRequest(req)
	case *uamsg.ReadRequest:
		rsp, err = secChan.handleReadRequest(req)
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
		length, err := secChan.conn.Write(content)
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
