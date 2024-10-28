package enc

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDecodeAcknowledgeMessage(t *testing.T) {
	bs := []byte{0x41, 0x43, 0x4b, 0x46, 0x1c, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0x0, 0x0, 0xff, 0xff, 0x0, 0x0, 0x0, 0x0, 0x20, 0x0, 0x40, 0x0, 0x0, 0x0}
	decoder := NewDefaultDecoder(bytes.NewBuffer(bs), 10240)
	msg, err := decoder.ReadMsg()
	require.NoError(t, err)
	require.NotNil(t, msg)
}
