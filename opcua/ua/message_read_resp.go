package ua

import (
	"github.com/libgox/buffer"
)

type MessageReadResp struct {
}

func DecodeMessageReadResp(buf *buffer.Buffer) (msg *MessageReadResp, err error) {
	msg = &MessageReadResp{}
	return msg, nil
}

func (m *MessageReadResp) Length() int {
	length := 0
	length += LenMessageType
	length += LenChunkType
	length += LenMessageSize
	return length
}

func (m *MessageReadResp) Buffer() (*buffer.Buffer, error) {
	buf := buffer.NewBuffer(m.Length())
	if _, err := buf.Write([]byte{'M', 'S', 'G'}); err != nil {
		return nil, err
	}
	if _, err := buf.Write([]byte{'F'}); err != nil {
		return nil, err
	}
	return buf, nil
}
