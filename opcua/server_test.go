package opcua

import (
	"log/slog"
	"testing"

	"github.com/protocol-laboratory/opcua-go/opcua/ua"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testSimpleClientServer struct {
	client *Client
	server *Server
}

func newTestSimpleClientServer(t *testing.T, serverConfig *ServerConfig, clientConfig *ClientConfig) *testSimpleClientServer {
	testLogger := slog.Default()

	serverConfig.Host = "localhost"
	serverConfig.Port = 0
	serverConfig.Logger = testLogger
	server, err := NewServer(serverConfig)
	require.NoError(t, err, "Server creation should not fail")
	port, err := server.Run()
	require.NoError(t, err, "Server should start without error")

	clientConfig.Logger = testLogger
	clientConfig.Address = Address{
		Host: serverConfig.Host,
		Port: port,
	}
	client, err := NewClient(clientConfig)
	require.NoError(t, err, "Client creation should not fail")

	return &testSimpleClientServer{
		client: client,
		server: server,
	}
}

func (s *testSimpleClientServer) close() {
	if s.client != nil {
		s.client.close()
	}
	if s.server != nil {
		_ = s.server.Close()
	}
}

func TestStartWithZeroPort(t *testing.T) {
	config := &ServerConfig{
		Host:   "localhost",
		Port:   0,
		Logger: slog.Default(),
	}

	server, err := NewServer(config)
	require.NoError(t, err)

	port, err := server.Run()
	require.NoError(t, err, "Server should start without error")

	assert.Greater(t, port, 0, "Expected a valid port to be assigned, but got %d", port)

	err = server.Close()
	assert.NoError(t, err, "Server should close without error")
}

func TestClientConnect(t *testing.T) {
	serverConfig := &ServerConfig{}
	clientConfig := &ClientConfig{}

	simpleClientServer := newTestSimpleClientServer(t, serverConfig, clientConfig)
	defer simpleClientServer.close()
}

func TestClientMessageHello(t *testing.T) {
	serverConfig := &ServerConfig{}
	clientConfig := &ClientConfig{}

	simpleClientServer := newTestSimpleClientServer(t, serverConfig, clientConfig)
	defer simpleClientServer.close()

	messageAcknowledge, err := simpleClientServer.client.Hello(&ua.MessageHello{
		Version:           0,
		ReceiveBufferSize: 65535,
		SendBufferSize:    65535,
		MaxMessageSize:    2097152,
		MaxChunkCount:     0,
		EndpointUrl:       "opc.tcp://localhost:4840/opcua",
	})
	require.NoError(t, err)

	assert.Equal(t, uint32(0), messageAcknowledge.Version)
	assert.Equal(t, uint32(65535), messageAcknowledge.ReceiveBufferSize)
	assert.Equal(t, uint32(65535), messageAcknowledge.SendBufferSize)
	assert.Equal(t, uint32(2097152), messageAcknowledge.MaxMessageSize)
	assert.Equal(t, uint32(64), messageAcknowledge.MaxChunkCount)
}
