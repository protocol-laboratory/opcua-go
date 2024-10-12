package ua

import (
	"github.com/libgox/buffer"
)

type MessageCloseSecureChannel struct {
	SecureChannelId uint32
}

func DecodeMessageCloseSecureChannel(buf *buffer.Buffer) (msg *MessageCloseSecureChannel, err error) {
	msg = &MessageCloseSecureChannel{}
	return msg, nil
}

func (m *MessageCloseSecureChannel) Length() int {
	length := 0
	length += LenMessageType
	length += LenChunkType
	length += LenMessageSize
	return length
}

func (m *MessageCloseSecureChannel) Buffer() (*buffer.Buffer, error) {
	buf := buffer.NewBuffer(m.Length())
	if _, err := buf.Write([]byte{'C', 'L', 'O'}); err != nil {
		return nil, err
	}
	if _, err := buf.Write([]byte{'F'}); err != nil {
		return nil, err
	}
	if err := buf.WriteUInt32Le(uint32(m.Length())); err != nil {
		return nil, err
	}
	return buf, nil
}
