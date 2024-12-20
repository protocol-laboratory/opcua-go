package ua

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeMessageBrowseResp(t *testing.T) {
	buffer := hex2Buffer(t, "4d534746480000000200000001000000fb030000c500000001001202b0c4be3815fcda01c50000000000000000ffffffff0000000100000000000000ffffffff00000000ffffffff")
	err := buffer.Skip(8)
	require.Nil(t, err)
}

func TestEncodeMessageBrowseResp(t *testing.T) {
	msg := &MessageBrowseResp{}
	_, err := msg.Buffer()
	require.Nil(t, err)
}
