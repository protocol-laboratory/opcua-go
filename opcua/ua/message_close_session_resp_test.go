package ua

import (
	"github.com/shoothzj/gox/testx"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDecodeMessageCloseSessionResp(t *testing.T) {
	buffer := testx.Hex2Buffer(t, "1e00000060090900005406400000000000000000000000000000000100000000000000000000000000000001d11afd7da59b38d0e146c88b801816eb005c00000101080a1f86fec7f584f6454d5347463400000002000000010000001c040000e60000000100dc01c07f353d15fcda01e60000000000000000ffffffff000000")
	err := buffer.Skip(8)
	require.Nil(t, err)
}

func TestEncodeMessageCloseSessionResp(t *testing.T) {
	msg := &MessageCloseSessionResp{}
	_, err := msg.Buffer()
	require.Nil(t, err)
}
