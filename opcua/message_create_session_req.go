package opcua

import "github.com/shoothzj/gox/buffer"

type MessageCreateSessionReq struct {
}

func DecodeMessageCreateSessionReq(buf *buffer.Buffer) (msg *MessageCreateSessionReq, err error) {
	msg = &MessageCreateSessionReq{}
	return msg, nil
}

func (m *MessageCreateSessionReq) Length() int {
	length := 0
	length += LenMessageType
	length += LenChunkType
	length += LenMessageSize
	return length
}

func (m *MessageCreateSessionReq) Buffer() (*buffer.Buffer, error) {
	buf := buffer.NewBuffer(m.Length())
	if _, err := buf.Write([]byte{'M', 'S', 'G'}); err != nil {
		return nil, err
	}
	if _, err := buf.Write([]byte{'F'}); err != nil {
		return nil, err
	}
	return buf, nil
}
