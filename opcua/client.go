package opcua

import "log/slog"

type ClientConfig struct {
	logger *slog.Logger
}

type Client struct {
	logger *slog.Logger
}

func NewClient(config *ClientConfig) *Client {
	client := &Client{
		logger: config.logger,
	}
	return client
}
