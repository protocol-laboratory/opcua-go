package opcua

import (
	"errors"
	"log/slog"

	"github.com/protocol-laboratory/opcua-go/opcua/enc"
	"github.com/protocol-laboratory/opcua-go/opcua/uamsg"
)

type SecureChannel struct {
	conn      *opcuaConn
	channelId uint32
	logger    *slog.Logger

	receiveMaxChunkSize    uint32
	sendMaxChunkSize       uint32
	maxChunkCount          uint32
	maxResponseMessageSize uint32
	endpointUrl            string

	// TODO conn set read timeout
	decoder enc.Decoder
	encoder enc.Encoder
}

const (
	// ProtocolVersion https://reference.opcfoundation.org/Core/Part6/v105/docs/7.1.2.3
	ProtocolVersion uint32 = 0
)

func newSecureChannel(conn *opcuaConn, svcConf *ServerConfig, channelId uint32, logger *slog.Logger) *SecureChannel {
	return &SecureChannel{
		conn:      conn,
		channelId: channelId,
		logger:    logger,
		decoder:   enc.NewDefaultDecoder(conn.conn, int64(svcConf.ReceiverBufferSize)),
		encoder:   enc.NewDefaultEncoder(),
	}
}

func (secChan *SecureChannel) open() error {
	secChan.logger.Info("handling hello message")
	err := secChan.handleHello()
	if err != nil {
		return err
	}

	secChan.logger.Info("handling open secure channel message")
	err = secChan.serve()
	if err != nil {
		secChan.logger.Error("serve error", slog.Any("err", err))
		return err
	}
	return nil
}

func (secChan *SecureChannel) handleHello() error {
	req, err := secChan.decoder.ReadMsg()
	if err != nil {
		return err
	}

	if req.MessageType != uamsg.HelloMessageType {
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
	return nil
}
