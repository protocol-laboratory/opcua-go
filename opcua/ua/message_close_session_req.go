package ua

import (
	"github.com/libgox/buffer"
)

type MessageCloseSessionReq struct {
}

func DecodeMessageCloseSessionReq(buf *buffer.Buffer) (msg *MessageCloseSessionReq, err error) {
	msg = &MessageCloseSessionReq{}
	return msg, nil
}

func (m *MessageCloseSessionReq) Length() int {
	length := 0
	length += LenMessageType
	length += LenChunkType
	length += LenMessageSize
	return length
}

func (m *MessageCloseSessionReq) Buffer() (*buffer.Buffer, error) {
	buf := buffer.NewBuffer(m.Length())
	if _, err := buf.Write([]byte{'M', 'S', 'G'}); err != nil {
		return nil, err
	}
	if _, err := buf.Write([]byte{'F'}); err != nil {
		return nil, err
	}
	return buf, nil
}
