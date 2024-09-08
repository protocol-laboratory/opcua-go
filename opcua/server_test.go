package opcua

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log/slog"
	"testing"
)

func TestStartWithZeroPort(t *testing.T) {
	config := &ServerConfig{
		Host:   "localhost",
		Port:   0,
		logger: slog.Default(),
	}

	server := NewServer(config)

	port, err := server.Run()
	require.NoError(t, err, "Server should start without error")

	assert.Greater(t, port, 0, "Expected a valid port to be assigned, but got %d", port)

	err = server.Close()
	assert.NoError(t, err, "Server should close without error")
}
