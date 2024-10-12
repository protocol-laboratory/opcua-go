package ua

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeMessageAcknowledge(t *testing.T) {
	buffer := hex2Buffer(t, "41434b461c00000000000000ffff0000ffff00000000200040000000")
	err := buffer.Skip(8)
	require.Nil(t, err)
	msg, err := DecodeMessageAcknowledge(buffer)
	require.Nil(t, err)
	require.NotNil(t, msg)
	assert.Equal(t, uint32(0), msg.Version)
	assert.Equal(t, uint32(65535), msg.ReceiveBufferSize)
	assert.Equal(t, uint32(65535), msg.SendBufferSize)
	assert.Equal(t, uint32(2097152), msg.MaxMessageSize)
	assert.Equal(t, uint32(64), msg.MaxChunkCount)
}

func TestEncodeMessageAcknowledge(t *testing.T) {
	msg := &MessageAcknowledge{
		Version:           0,
		ReceiveBufferSize: 65535,
		SendBufferSize:    65535,
		MaxMessageSize:    2097152,
		MaxChunkCount:     64,
	}
	buffer, err := msg.Buffer()
	require.Nil(t, err)
	assert.Equal(t, hex2Buffer(t, "41434b461c00000000000000ffff0000ffff00000000200040000000"), buffer)
}
