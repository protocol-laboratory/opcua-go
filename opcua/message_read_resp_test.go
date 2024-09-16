package opcua

import (
	"github.com/shoothzj/gox/testx"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDecodeMessageReadResp(t *testing.T) {
	buffer := testx.Hex2Buffer(t, "4d5347464a00000002000000010000001a040000e400000001007a02207b4e3c15fcda01e40000000000000000ffffffff00000001000000050600000000206b273c15fcda01ffffffff")
	err := buffer.Skip(8)
	require.Nil(t, err)
}

func TestEncodeMessageReadResp(t *testing.T) {
	msg := &MessageReadResp{}
	_, err := msg.Buffer()
	require.Nil(t, err)
}
