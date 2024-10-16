package opcua

import (
	"github.com/protocol-laboratory/opcua-go/opcua/uamsg"
)

func (secChan *SecureChannel) handleOpenSecureChannelRequest(req *uamsg.Message) (*uamsg.Message, error) {
	openSecureChannelRequest, ok := req.MessageBody.(*uamsg.OpenSecureChannelRequest)
	if !ok {
		return nil, ErrInvalidMessageBody
	}
	err := secChan.handler.BeforeOpenSecureChannel(secChan.conn, openSecureChannelRequest)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (secChan *SecureChannel) handleCloseSecureChannelRequest(req *uamsg.Message) (*uamsg.Message, error) {
	closeSecureChannelRequest, ok := req.MessageBody.(*uamsg.CloseSecureChannelRequest)
	if !ok {
		return nil, ErrInvalidMessageBody
	}
	err := secChan.handler.BeforeCloseSecureChannel(secChan.conn, closeSecureChannelRequest)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (secChan *SecureChannel) handleCreateSessionRequest(req *uamsg.Message) (*uamsg.Message, error) {
	createSessionRequest, ok := req.MessageBody.(*uamsg.CreateSessionRequest)
	if !ok {
		return nil, ErrInvalidMessageBody
	}
	err := secChan.handler.BeforeCreateSession(secChan.conn, createSessionRequest)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (secChan *SecureChannel) handleActivateSessionRequest(req *uamsg.Message) (*uamsg.Message, error) {
	activateSessionRequest, ok := req.MessageBody.(*uamsg.ActivateSessionRequest)
	if !ok {
		return nil, ErrInvalidMessageBody
	}
	err := secChan.handler.BeforeActivateSession(secChan.conn, activateSessionRequest)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (secChan *SecureChannel) handleCloseSessionRequest(req *uamsg.Message) (*uamsg.Message, error) {
	_, ok := req.MessageBody.(*uamsg.CloseSessionRequest)
	if !ok {
		return nil, ErrInvalidMessageBody
	}
	return nil, nil
}

func (secChan *SecureChannel) handleGetEndpoints(req *uamsg.Message) (*uamsg.Message, error) {
	getEndpointsRequest, ok := req.MessageBody.(*uamsg.GetEndpointsRequest)
	if !ok {
		return nil, ErrInvalidMessageBody
	}
	err := secChan.handler.BeforeGetEndpoints(secChan.conn, getEndpointsRequest)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (secChan *SecureChannel) handleBrowseRequest(req *uamsg.Message) (*uamsg.Message, error) {
	browseRequest, ok := req.MessageBody.(*uamsg.BrowseRequest)
	if !ok {
		return nil, ErrInvalidMessageBody
	}
	err := secChan.handler.BeforeBrowse(secChan.conn, browseRequest)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (secChan *SecureChannel) handleReadRequest(req *uamsg.Message) (*uamsg.Message, error) {
	readRequest, ok := req.MessageBody.(*uamsg.ReadRequest)
	if !ok {
		return nil, ErrInvalidMessageBody
	}
	err := secChan.handler.BeforeRead(secChan.conn, readRequest)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
