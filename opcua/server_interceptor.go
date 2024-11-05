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

	AfterHello(conn *Conn, helloMessage *uamsg.HelloMessageExtras) error
	AfterOpenSecureChannel(conn *Conn, req *uamsg.OpenSecureChannelRequest, resp *uamsg.OpenSecureChannelResponse) error
	AfterCloseSecureChannel(conn *Conn, req *uamsg.CloseSecureChannelRequest, resp *uamsg.CloseSecureChannelResponse) error
	AfterCreateSession(conn *Conn, req *uamsg.CreateSessionRequest, resp *uamsg.CreateSessionResponse) error
	AfterActivateSession(conn *Conn, req *uamsg.ActivateSessionRequest, resp *uamsg.ActivateSessionResponse) error
	AfterCloseSession(conn *Conn, req *uamsg.CloseSessionRequest, resp *uamsg.CloseSessionResponse) error
	AfterGetEndpoints(conn *Conn, req *uamsg.GetEndpointsRequest, resp *uamsg.GetEndpointsResponse) error
	AfterBrowse(conn *Conn, req *uamsg.BrowseRequest, resp *uamsg.BrowseResponse) error
	AfterRead(conn *Conn, req *uamsg.ReadRequest, resp *uamsg.ReadResponse) error
}

var _ ServerInterceptor = (*NoopServerInterceptor)(nil)

type NoopServerInterceptor struct {
}

func (n NoopServerInterceptor) ConnectionOpened(conn *Conn) {
}

func (n NoopServerInterceptor) ConnectionClosed(conn *Conn) {
}

func (n NoopServerInterceptor) BeforeHello(conn *Conn, helloMessage *uamsg.HelloMessageExtras) error {
	return nil
}

func (n NoopServerInterceptor) BeforeOpenSecureChannel(conn *Conn, openSecureChannelMessage *uamsg.OpenSecureChannelRequest) error {
	return nil
}

func (n NoopServerInterceptor) BeforeCloseSecureChannel(conn *Conn, closeSecureChannelMessage *uamsg.CloseSecureChannelRequest) error {
	return nil
}

func (n NoopServerInterceptor) BeforeCreateSession(conn *Conn, createSessionMessage *uamsg.CreateSessionRequest) error {
	return nil
}

func (n NoopServerInterceptor) BeforeActivateSession(conn *Conn, activateSessionMessage *uamsg.ActivateSessionRequest) error {
	return nil
}

func (n NoopServerInterceptor) BeforeCloseSession(conn *Conn, closeSessionMessage *uamsg.CloseSessionRequest) error {
	return nil
}

func (n NoopServerInterceptor) BeforeGetEndpoints(conn *Conn, getEndpointsMessage *uamsg.GetEndpointsRequest) error {
	return nil
}

func (n NoopServerInterceptor) BeforeBrowse(conn *Conn, browseMessage *uamsg.BrowseRequest) error {
	return nil
}

func (n NoopServerInterceptor) BeforeRead(conn *Conn, readMessage *uamsg.ReadRequest) error {
	return nil
}

func (n NoopServerInterceptor) AfterHello(conn *Conn, helloMessage *uamsg.HelloMessageExtras) error {
	return nil
}

func (n NoopServerInterceptor) AfterOpenSecureChannel(conn *Conn, req *uamsg.OpenSecureChannelRequest, resp *uamsg.OpenSecureChannelResponse) error {
	return nil
}

func (n NoopServerInterceptor) AfterCloseSecureChannel(conn *Conn, req *uamsg.CloseSecureChannelRequest, resp *uamsg.CloseSecureChannelResponse) error {
	return nil
}

func (n NoopServerInterceptor) AfterCreateSession(conn *Conn, req *uamsg.CreateSessionRequest, resp *uamsg.CreateSessionResponse) error {
	return nil
}

func (n NoopServerInterceptor) AfterActivateSession(conn *Conn, req *uamsg.ActivateSessionRequest, resp *uamsg.ActivateSessionResponse) error {
	return nil
}

func (n NoopServerInterceptor) AfterCloseSession(conn *Conn, req *uamsg.CloseSessionRequest, resp *uamsg.CloseSessionResponse) error {
	return nil
}

func (n NoopServerInterceptor) AfterGetEndpoints(conn *Conn, req *uamsg.GetEndpointsRequest, resp *uamsg.GetEndpointsResponse) error {
	return nil
}

func (n NoopServerInterceptor) AfterBrowse(conn *Conn, req *uamsg.BrowseRequest, resp *uamsg.BrowseResponse) error {
	return nil
}

func (n NoopServerInterceptor) AfterRead(conn *Conn, req *uamsg.ReadRequest, resp *uamsg.ReadResponse) error {
	return nil
}
