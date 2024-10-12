package ua

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeMessageReadReq(t *testing.T) {
	buffer := hex2Buffer(t, "4d534746800000000200000001000000e4000000e4000000010077020500002000000048cce1823313199f88bb583a480954a576d4fae160deef9bf25daa6e50bfbc39fa284e3c15fcda01e400000000000000ffffffffa00f0000000000000000000000000000000000010000000100d3080d000000ffffffff0000ffffffff")
	err := buffer.Skip(8)
	require.Nil(t, err)
}

func TestEncodeMessageReadReq(t *testing.T) {
	msg := &MessageReadReq{}
	_, err := msg.Buffer()
	require.Nil(t, err)
}
