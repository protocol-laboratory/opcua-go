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

type NoopServerHandler struct {
}

func (d *NoopServerHandler) ConnectionOpened(conn *Conn) {
}

func (d *NoopServerHandler) ConnectionClosed(conn *Conn) {
}

func (d *NoopServerHandler) BeforeHello(conn *Conn, helloMessage *uamsg.HelloMessageExtras) error {
	return nil
}

func (d *NoopServerHandler) BeforeOpenSecureChannel(conn *Conn, openSecureChannelMessage *uamsg.OpenSecureChannelRequest) error {
	return nil
}

func (d *NoopServerHandler) BeforeCloseSecureChannel(conn *Conn, closeSecureChannelMessage *uamsg.CloseSecureChannelRequest) error {
	return nil
}

func (d *NoopServerHandler) BeforeCreateSession(conn *Conn, createSessionMessage *uamsg.CreateSessionRequest) error {
	return nil
}

func (d *NoopServerHandler) BeforeActivateSession(conn *Conn, activateSessionMessage *uamsg.ActivateSessionRequest) error {
	return nil
}

func (d *NoopServerHandler) BeforeCloseSession(conn *Conn, closeSessionMessage *uamsg.CloseSessionRequest) error {
	return nil
}

func (d *NoopServerHandler) BeforeGetEndpoints(conn *Conn, getEndpointsMessage *uamsg.GetEndpointsRequest) error {
	return nil
}

func (d *NoopServerHandler) BeforeBrowse(conn *Conn, browseMessage *uamsg.BrowseRequest) error {
	return nil
}

func (d *NoopServerHandler) BeforeRead(conn *Conn, readMessage *uamsg.ReadRequest) error {
	return nil
}
