package opcua

import "github.com/protocol-laboratory/opcua-go/opcua/uamsg"

type ServerHandler interface {
	ConnectionOpened(conn *Conn)
	ConnectionClosed(conn *Conn)

	BeforeHello(conn *Conn, helloMessage *uamsg.HelloMessageExtras) error
	BeforeOpenSecureChannel(conn *Conn, openSecureChannelMessage *uamsg.OpenSecureChannelRequest) error
	BeforeCloseSecureChannel(conn *Conn, closeSecureChannelMessage *uamsg.CloseSecureChannelRequest) error
	BeforeCreateSession(conn *Conn, createSessionMessage *uamsg.CreateSessionRequest) error
	BeforeActivateSession(conn *Conn, activateSessionMessage *uamsg.ActivateSessionRequest) error
	BeforeCloseSession(conn *Conn, closeSessionMessage *uamsg.CloseSessionRequest) error
	BeforeGetEndpoints(conn *Conn, getEndpointsMessage *uamsg.GetEndpointsRequest) error
	BeforeBrowse(conn *Conn, browseMessage *uamsg.BrowseRequest) error
	BeforeRead(conn *Conn, readMessage *uamsg.ReadRequest) error
}

type DefaultServerHandler struct {
}

func (d *DefaultServerHandler) ConnectionOpened(conn *Conn) {
}

func (d *DefaultServerHandler) ConnectionClosed(conn *Conn) {
}

func (d *DefaultServerHandler) BeforeHello(conn *Conn, helloMessage *uamsg.HelloMessageExtras) error {
	return nil
}

func (d *DefaultServerHandler) BeforeOpenSecureChannel(conn *Conn, openSecureChannelMessage *uamsg.OpenSecureChannelRequest) error {
	return nil
}

func (d *DefaultServerHandler) BeforeCloseSecureChannel(conn *Conn, closeSecureChannelMessage *uamsg.CloseSecureChannelRequest) error {
	return nil
}

func (d *DefaultServerHandler) BeforeCreateSession(conn *Conn, createSessionMessage *uamsg.CreateSessionRequest) error {
	return nil
}

func (d *DefaultServerHandler) BeforeActivateSession(conn *Conn, activateSessionMessage *uamsg.ActivateSessionRequest) error {
	return nil
}

func (d *DefaultServerHandler) BeforeCloseSession(conn *Conn, closeSessionMessage *uamsg.CloseSessionRequest) error {
	return nil
}

func (d *DefaultServerHandler) BeforeGetEndpoints(conn *Conn, getEndpointsMessage *uamsg.GetEndpointsRequest) error {
	return nil
}

func (d *DefaultServerHandler) BeforeBrowse(conn *Conn, browseMessage *uamsg.BrowseRequest) error {
	return nil
}

func (d *DefaultServerHandler) BeforeRead(conn *Conn, readMessage *uamsg.ReadRequest) error {
	return nil
}
