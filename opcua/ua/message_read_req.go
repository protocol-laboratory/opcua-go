package ua

import (
	"github.com/shoothzj/gox/buffer"
)

type MessageReadReq struct {
}

func DecodeMessageReadReq(buf *buffer.Buffer) (msg *MessageReadReq, err error) {
	msg = &MessageReadReq{}
	return msg, nil
}

func (m *MessageReadReq) Length() int {
	length := 0
	length += LenMessageType
	length += LenChunkType
	length += LenMessageSize
	return length
}

func (m *MessageReadReq) Buffer() (*buffer.Buffer, error) {
	buf := buffer.NewBuffer(m.Length())
	if _, err := buf.Write([]byte{'M', 'S', 'G'}); err != nil {
		return nil, err
	}
	if _, err := buf.Write([]byte{'F'}); err != nil {
		return nil, err
	}
	return buf, nil
}
