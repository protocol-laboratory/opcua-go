package enc

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeOpenSecureChannelMessage(t *testing.T) {
	bs := []byte{0x4f, 0x50, 0x4e, 0x46, 0x87, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x2f, 0x0, 0x0, 0x0, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6f, 0x70, 0x63, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x55, 0x41, 0x2f, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x23, 0x4e, 0x6f, 0x6e, 0x65, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x1, 0x0, 0xc1, 0x1, 0x80, 0xca, 0xa9, 0xcb, 0xf7, 0x80, 0xd9, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x70, 0xa3, 0xa9, 0xcb, 0xf7, 0x80, 0xd9, 0x1, 0x80, 0xee, 0x36, 0x0, 0xff, 0xff, 0xff, 0xff}
	decoder := NewDefaultDecoder(bytes.NewBuffer(bs), 10240)
	msg, err := decoder.ReadMsg()
	require.NoError(t, err)
	require.NotNil(t, msg)
}
