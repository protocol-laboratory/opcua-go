package enc

import "opcua-go/opcua/uamsg"

type Decoder interface {
	SetMaxBufferSize(size int)
	SetMessageRsvCh(msgCh chan<- *uamsg.Message)
	Decode(buf []byte) error
}
