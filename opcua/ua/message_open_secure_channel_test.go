package ua

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeMessageOpenSecureChannel(t *testing.T) {
	buffer := hex2Buffer(t, "4f504e4684000000000000002f000000687474703a2f2f6f7063666f756e646174696f6e2e6f72672f55412f5365637572697479506f6c696379234e6f6e65ffffffffffffffff01000000010000000100be0100000092a8cbf780d9010000000000000000ffffffff60ea0000000000000000000000000001000000ffffffff80ee3600")
	err := buffer.Skip(8)
	require.Nil(t, err)
}

func TestEncodeMessageOpenSecureChannel(t *testing.T) {
	msg := &MessageOpenSecureChannel{}
	_, err := msg.Buffer()
	require.Nil(t, err)
}
