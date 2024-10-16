package opcua

import (
	"net"
)

type Conn struct {
	net.Conn
	context any
}

func (c *Conn) SetContext(value any) {
	c.context = value
}

func (c *Conn) Context() (value any) {
	return c.context
}

func (c *Conn) NetConn() net.Conn {
	return c.Conn
}

func (c *Conn) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Conn) LocalAddr() net.Addr {
	return c.Conn.LocalAddr()
}
