package ua

import (
	"github.com/libgox/buffer"
)

type MessageCloseSessionResp struct {
}

func DecodeMessageCloseSessionResp(buf *buffer.Buffer) (msg *MessageCloseSessionResp, err error) {
	msg = &MessageCloseSessionResp{}
	return msg, nil
}

func (m *MessageCloseSessionResp) Length() int {
	length := 0
	length += LenMessageType
	length += LenChunkType
	length += LenMessageSize
	return length
}

func (m *MessageCloseSessionResp) Buffer() (*buffer.Buffer, error) {
	buf := buffer.NewBuffer(m.Length())
	if _, err := buf.Write([]byte{'M', 'S', 'G'}); err != nil {
		return nil, err
	}
	if _, err := buf.Write([]byte{'F'}); err != nil {
		return nil, err
	}
	return buf, nil
}
