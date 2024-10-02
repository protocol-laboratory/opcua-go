package ua

import (
	"github.com/shoothzj/gox/buffer"
)

type MessageGetEndpointsResp struct {
}

func DecodeMessageGetEndpointsResp(buf *buffer.Buffer) (msg *MessageGetEndpointsResp, err error) {
	msg = &MessageGetEndpointsResp{}
	return msg, nil
}

func (m *MessageGetEndpointsResp) Length() int {
	length := 0
	length += LenMessageType
	length += LenChunkType
	length += LenMessageSize
	return length
}

func (m *MessageGetEndpointsResp) Buffer() (*buffer.Buffer, error) {
	buf := buffer.NewBuffer(m.Length())
	if _, err := buf.Write([]byte{'M', 'S', 'G'}); err != nil {
		return nil, err
	}
	if _, err := buf.Write([]byte{'F'}); err != nil {
		return nil, err
	}
	return buf, nil
}
