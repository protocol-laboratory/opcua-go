package opcua

import (
	"github.com/shoothzj/gox/testx"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDecodeMessageBrowseReq(t *testing.T) {
	buffer := testx.Hex2Buffer(t, "4d534746870000000200000001000000c8000000c800000001000f020500002000000048cce1823313199f88bb583a480954a576d4fae160deef9bf25daa6e50bfbc39f6c9be3815fcda01c800000000000000ffffffffa00f00000000000000000000000000000000000000000000000100000001000f3d00000000002d01000000003f000000")
	err := buffer.Skip(8)
	require.Nil(t, err)
}

func TestEncodeMessageBrowseReq(t *testing.T) {
	msg := &MessageBrowseReq{}
	_, err := msg.Buffer()
	require.Nil(t, err)
}
