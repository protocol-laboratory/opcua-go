package ua

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeMessageCloseSecureChannel(t *testing.T) {
	buffer := hex2Buffer(t, "4d5347465f0000000200000001000000e6000000e60000000100d9010500002000000048cce1823313199f88bb583a480954a576d4fae160deef9bf25daa6e50bfbc396ee1343d15fcda01e600000000000000ffffffffa00f000000000001")
	err := buffer.Skip(8)
	require.Nil(t, err)
}

func TestEncodeMessageCloseSecureChannel(t *testing.T) {
	msg := &MessageCloseSecureChannel{}
	_, err := msg.Buffer()
	require.Nil(t, err)
}
