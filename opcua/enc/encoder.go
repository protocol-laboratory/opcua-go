package enc

import "opcua-go/opcua/uamsg"

type Encoder interface {
	Encode(v *uamsg.Message, chunksize int) ([][]byte, error)
}
