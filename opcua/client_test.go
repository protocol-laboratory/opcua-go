package opcua

import (
	"github.com/shoothzj/gox/netx"
	"github.com/stretchr/testify/require"
	"log/slog"
	"testing"
)

func TestClientConnect(t *testing.T) {
	logger := slog.Default()

	serverConfig := &ServerConfig{
		Host:   "localhost",
		Port:   0,
		Logger: logger,
	}
	server := NewServer(serverConfig)

	port, err := server.Run()
	require.NoError(t, err)
	defer server.Close()

	clientConfig := &ClientConfig{
		Address: netx.Address{
			Host: "127.0.0.1",
			Port: port,
		},
		Logger: logger,
	}
	client, err := NewClient(clientConfig)
	require.NoError(t, err)

	client.close()
}
