package ua

import (
	"github.com/shoothzj/gox/testx"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDecodeMessageGetEndpointsReq(t *testing.T) {
	buffer := testx.Hex2Buffer(t, "4d534746a8000000010000000100000002000000020000000100ac010000a018aacbf780d9010000000000000000ffffffff60ea00000000001e0000006f70632e7463703a2f2f6c6f63616c686f73743a31323638362f6d696c6f000000000100000041000000687474703a2f2f6f7063666f756e646174696f6e2e6f72672f55412d50726f66696c652f5472616e73706f72742f75617463702d756173632d756162696e617279")
	err := buffer.Skip(8)
	require.Nil(t, err)
}

func TestEncodeMessageGetEndpointsReq(t *testing.T) {
	msg := &MessageGetEndpointsReq{}
	_, err := msg.Buffer()
	require.Nil(t, err)
}
