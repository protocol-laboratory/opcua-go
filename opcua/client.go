package opcua

import "log/slog"

type ClientConfig struct {
	logger *slog.Logger
}

type Client struct {
	config *ClientConfig
	logger *slog.Logger
}

func NewClient(config *ClientConfig) *Client {
	client := &Client{
		config: config,
		logger: config.logger,
	}
	return client
}
