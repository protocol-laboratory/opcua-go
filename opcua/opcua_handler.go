package opcua

import "net"

type ServerHandler interface {
	ConnectionOpened(conn net.Conn)
	ConnectionClosed(conn net.Conn)
}
