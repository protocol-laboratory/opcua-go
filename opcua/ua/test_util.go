package ua

import (
	"encoding/hex"
	"testing"

	"github.com/libgox/buffer"
	"github.com/stretchr/testify/require"
)

func hex2Buffer(t *testing.T, str string) *buffer.Buffer {
	bytes, err := hex.DecodeString(str)
	require.NoError(t, err)
	return buffer.NewBufferFromBytes(bytes)
}
