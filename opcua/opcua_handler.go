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

	ReadDataValue(conn *Conn, nodeToRead *uamsg.ReadValueId) (*uamsg.Variant, error)
}

var _ ServerHandler = (*NoopServerHandler)(nil)

type NoopServerHandler struct {
}

func (n *NoopServerHandler) ConnectionOpened(conn *Conn) {
}

func (n *NoopServerHandler) ConnectionClosed(conn *Conn) {
}

func (n *NoopServerHandler) BeforeHello(conn *Conn, helloMessage *uamsg.HelloMessageExtras) error {
	return nil
}

func (n *NoopServerHandler) BeforeOpenSecureChannel(conn *Conn, openSecureChannelMessage *uamsg.OpenSecureChannelRequest) error {
	return nil
}

func (n *NoopServerHandler) BeforeCloseSecureChannel(conn *Conn, closeSecureChannelMessage *uamsg.CloseSecureChannelRequest) error {
	return nil
}

func (n *NoopServerHandler) BeforeCreateSession(conn *Conn, createSessionMessage *uamsg.CreateSessionRequest) error {
	return nil
}

func (n *NoopServerHandler) BeforeActivateSession(conn *Conn, activateSessionMessage *uamsg.ActivateSessionRequest) error {
	return nil
}

func (n *NoopServerHandler) BeforeCloseSession(conn *Conn, closeSessionMessage *uamsg.CloseSessionRequest) error {
	return nil
}

func (n *NoopServerHandler) BeforeGetEndpoints(conn *Conn, getEndpointsMessage *uamsg.GetEndpointsRequest) error {
	return nil
}

func (n *NoopServerHandler) BeforeBrowse(conn *Conn, browseMessage *uamsg.BrowseRequest) error {
	return nil
}

func (n *NoopServerHandler) BeforeRead(conn *Conn, readMessage *uamsg.ReadRequest) error {
	return nil
}

func (n *NoopServerHandler) ReadDataValue(conn *Conn, nodeToRead *uamsg.ReadValueId) (*uamsg.Variant, error) {
	return nil, nil
}
