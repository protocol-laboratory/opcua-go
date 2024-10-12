package ua

import (
	"testing"

	"github.com/shoothzj/gox/testx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeMessageHello(t *testing.T) {
	buffer := testx.Hex2Buffer(t, "48454c463e00000000000000ffff0000ffff000000002000400000001e0000006f70632e7463703a2f2f6c6f63616c686f73743a31323638362f6d696c6f")
	err := buffer.Skip(8)
	require.Nil(t, err)
	msg, err := DecodeMessageHello(buffer)
	require.Nil(t, err)
	require.NotNil(t, msg)
	assert.Equal(t, uint32(0), msg.Version)
	assert.Equal(t, uint32(65535), msg.ReceiveBufferSize)
	assert.Equal(t, uint32(65535), msg.SendBufferSize)
	assert.Equal(t, uint32(2097152), msg.MaxMessageSize)
	assert.Equal(t, uint32(64), msg.MaxChunkCount)
	assert.Equal(t, "opc.tcp://localhost:12686/milo", msg.EndpointUrl)
}

func TestEncodeMessageHello(t *testing.T) {
	msg := &MessageHello{
		Version:           0,
		ReceiveBufferSize: 65535,
		SendBufferSize:    65535,
		MaxMessageSize:    2097152,
		MaxChunkCount:     64,
		EndpointUrl:       "opc.tcp://localhost:12686/milo",
	}
	buffer, err := msg.Buffer()
	require.Nil(t, err)
	assert.Equal(t, testx.Hex2Buffer(t, "48454c463e00000000000000ffff0000ffff000000002000400000001e0000006f70632e7463703a2f2f6c6f63616c686f73743a31323638362f6d696c6f"), buffer)
}
