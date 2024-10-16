package opcua

import "github.com/protocol-laboratory/opcua-go/opcua/uamsg"

type ServerHandler interface {
	ConnectionOpened(conn *Conn)
	ConnectionClosed(conn *Conn)

	ValidateHello(conn *Conn, helloMessage *uamsg.HelloMessageExtras) error
}

type DefaultServerHandler struct {
}

func (d DefaultServerHandler) ConnectionOpened(conn *Conn) {
}

func (d DefaultServerHandler) ConnectionClosed(conn *Conn) {
}

func (d DefaultServerHandler) ValidateHello(conn *Conn, helloMessage *uamsg.HelloMessageExtras) error {
	return nil
}
