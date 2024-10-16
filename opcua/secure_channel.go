package opcua

import "log/slog"

type SecureChannel struct {
	conn      *opcuaConn
	channelId uint32
	logger    *slog.Logger
}

func newSecureChannel(conn *opcuaConn, svcConf *ServerConfig, channelId uint32, logger *slog.Logger) *SecureChannel {
	return &SecureChannel{
		conn:      conn,
		channelId: channelId,
		logger:    logger,
	}
}
