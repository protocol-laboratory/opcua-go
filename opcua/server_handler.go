package opcua

import "github.com/protocol-laboratory/opcua-go/opcua/uamsg"

type ServerHandler interface {
	HandleRead(conn *Conn, readRequest *uamsg.ReadRequest) (*uamsg.ReadResponse, error)
}

var _ ServerHandler = (*NoopServerHandler)(nil)

type NoopServerHandler struct {
}

func (n NoopServerHandler) HandleRead(conn *Conn, readRequest *uamsg.ReadRequest) (*uamsg.ReadResponse, error) {
	return nil, nil
}
