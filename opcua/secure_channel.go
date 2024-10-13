package opcua

import (
	"errors"
	"sync/atomic"

	"github.com/protocol-laboratory/opcua-go/opcua/enc"
	"github.com/protocol-laboratory/opcua-go/opcua/uamsg"
	"github.com/protocol-laboratory/opcua-go/opcua/util"
)

var nextChannelId atomic.Uint32

type SecureChannel struct {
	conn *opcuaConn

	receiveMaxChunkSize    uint32
	sendMaxChunkSize       uint32
	maxChunkCount          uint32
	maxResponseMessageSize uint32
	endpointUrl            string

	decoder enc.Decoder
	//// TODO read timeout
	encoder enc.Encoder

	nextTokenId        atomic.Uint32
	nextSequenceNumber atomic.Uint32
}

const (
	ProtocolVersion uint32 = 0 // https://reference.opcfoundation.org/Core/Part6/v105/docs/7.1.2.3
)

func newSecureChannel(conn *opcuaConn, svcConf *ServerConfig) *SecureChannel {
	return &SecureChannel{
		conn:    conn,
		decoder: enc.NewDefaultDecoder(conn.conn, int64(svcConf.ReceiverBufferSize)),
		encoder: enc.NewDefaultEncoder(),
	}
}

func (secChan *SecureChannel) open() error {
	// the first message must be a Hello
	err := secChan.handleHello()
	if err != nil {
		return err
	}

	// the second message must be an OpenSecureChannel
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
		return errors.New("invalid message type")
	}

	helloBody, ok := req.MessageBody.(*uamsg.HelloMessageExtras)
	if !ok {
		return errors.New("invalid message type")
	}

	//// TODO should affect en/decoder's behavior
	secChan.receiveMaxChunkSize = helloBody.SendBufferSize // TODO limit
	secChan.sendMaxChunkSize = helloBody.ReceiveBufferSize // TODO limit
	secChan.maxChunkCount = helloBody.MaxChunkCount
	secChan.maxResponseMessageSize = helloBody.MaxMessageSize
	secChan.endpointUrl = helloBody.EndpointUrl

	rsp := &uamsg.Message{
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

	bytes, err := secChan.encoder.Encode(rsp, int(secChan.maxResponseMessageSize))
	if err != nil {
		return err
	}

	for _, content := range bytes {
		_, err = secChan.conn.conn.Write(content)
		if err != nil {
			return err
		}
	}

	//// TODO should return ERROR STATUS CODE to client if any error occur
	return nil
}

func (secChan *SecureChannel) handleOpenSecureChannel() error {
	req, err := secChan.decoder.ReadMsg()
	if err != nil {
		return err
	}

	if req.MessageHeader.MessageType != uamsg.OpenSecureChannelMessageType {
		return errors.New("invalid message type")
	}

	genericBody, ok := req.MessageBody.(*uamsg.GenericBody)
	if !ok {
		return errors.New("invalid message type")
	}

	openSecChanBody, ok := genericBody.Service.(*uamsg.OpenSecureChannelServiceRequest)
	if !ok {
		return errors.New("invalid message type")
	}

	//// TODO only support NONE mode yet
	if openSecChanBody.SecurityMode != uamsg.MessageSecurityModeNone {
		return errors.New("only support NONE security mode")
	}

	serverNonce := []byte{}
	channelId := getNextChannelId()
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
				NodeId: &uamsg.NodeId{
					EncodingType: uamsg.FourByte,
					Namespace:    0,
					Identifier:   uint16(449),
				},
			},
			Service: &uamsg.OpenSecureChannelServiceResponse{
				Header: &uamsg.ResponseHeader{
					Timestamp:     util.GetCurrentUaTimestamp(),
					RequestHandle: 0,
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

	bytes, err := secChan.encoder.Encode(rsp, int(secChan.maxResponseMessageSize))
	if err != nil {
		return err
	}

	for _, content := range bytes {
		_, err = secChan.conn.conn.Write(content)
		if err != nil {
			return err
		}
	}

	//// TODO should return ERROR STATUS CODE to client if any error occur
	return nil
}

func (secChan *SecureChannel) close() {
}

func (secChan *SecureChannel) fatal() {
}

func (secChan *SecureChannel) processRequests() error {
	return nil
}

func getNextChannelId() uint32 {
	return nextChannelId.Add(1)
}

func (secChan *SecureChannel) getNextTokenId() uint32 {
	return secChan.nextTokenId.Add(1)
}

func (secChan *SecureChannel) getNextSequenceNumber() uint32 {
	return secChan.nextSequenceNumber.Add(1)
}
