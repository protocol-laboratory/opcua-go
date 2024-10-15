package opcua

type SecureChannel struct {
	conn *opcuaConn
}

func newSecureChannel(conn *opcuaConn, svcConf *ServerConfig) *SecureChannel {
	return &SecureChannel{
		conn: conn,
	}
}
