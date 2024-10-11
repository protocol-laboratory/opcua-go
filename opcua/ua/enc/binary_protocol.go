package enc

import (
	"github.com/protocol-laboratory/opcua-go/opcua/ua"
)

type Encoder interface {
	Encode(msg ua.Message, chunkSize int) ([][]byte, error)
}

type Decoder interface {
	SetMaxBufferSize(size int)
	SetMessageReceiveChannel(ch chan<- ua.Message)
	Decode(buf []byte) error
}
