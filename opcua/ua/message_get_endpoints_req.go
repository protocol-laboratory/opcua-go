package ua

import (
	"github.com/shoothzj/gox/buffer"
)

type MessageGetEndpointsReq struct {
}

func DecodeMessageGetEndpointsReq(buf *buffer.Buffer) (msg *MessageGetEndpointsReq, err error) {
	msg = &MessageGetEndpointsReq{}
	return msg, nil
}

func (m *MessageGetEndpointsReq) Length() int {
	length := 0
	length += LenMessageType
	length += LenChunkType
	length += LenMessageSize
	return length
}

func (m *MessageGetEndpointsReq) Buffer() (*buffer.Buffer, error) {
	buf := buffer.NewBuffer(m.Length())
	if _, err := buf.Write([]byte{'M', 'S', 'G'}); err != nil {
		return nil, err
	}
	if _, err := buf.Write([]byte{'F'}); err != nil {
		return nil, err
	}
	return buf, nil
}
