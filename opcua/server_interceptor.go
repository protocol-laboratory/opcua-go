package opcua

import "github.com/protocol-laboratory/opcua-go/opcua/uamsg"

type ServerInterceptor interface {
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

var _ ServerInterceptor = (*NoopServerInterceptor)(nil)

type NoopServerInterceptor struct {
}

func (n *NoopServerInterceptor) ConnectionOpened(conn *Conn) {
}

func (n *NoopServerInterceptor) ConnectionClosed(conn *Conn) {
}

func (n *NoopServerInterceptor) BeforeHello(conn *Conn, helloMessage *uamsg.HelloMessageExtras) error {
	return nil
}

func (n *NoopServerInterceptor) BeforeOpenSecureChannel(conn *Conn, openSecureChannelMessage *uamsg.OpenSecureChannelRequest) error {
	return nil
}

func (n *NoopServerInterceptor) BeforeCloseSecureChannel(conn *Conn, closeSecureChannelMessage *uamsg.CloseSecureChannelRequest) error {
	return nil
}

func (n *NoopServerInterceptor) BeforeCreateSession(conn *Conn, createSessionMessage *uamsg.CreateSessionRequest) error {
	return nil
}

func (n *NoopServerInterceptor) BeforeActivateSession(conn *Conn, activateSessionMessage *uamsg.ActivateSessionRequest) error {
	return nil
}

func (n *NoopServerInterceptor) BeforeCloseSession(conn *Conn, closeSessionMessage *uamsg.CloseSessionRequest) error {
	return nil
}

func (n *NoopServerInterceptor) BeforeGetEndpoints(conn *Conn, getEndpointsMessage *uamsg.GetEndpointsRequest) error {
	return nil
}

func (n *NoopServerInterceptor) BeforeBrowse(conn *Conn, browseMessage *uamsg.BrowseRequest) error {
	return nil
}

func (n *NoopServerInterceptor) BeforeRead(conn *Conn, readMessage *uamsg.ReadRequest) error {
	return nil
}

func (n *NoopServerInterceptor) ReadDataValue(conn *Conn, nodeToRead *uamsg.ReadValueId) (*uamsg.Variant, error) {
	return nil, nil
}
