package opcua

import "github.com/shoothzj/gox/buffer"

type MessageBrowseReq struct {
}

func DecodeMessageBrowseReq(buf *buffer.Buffer) (msg *MessageBrowseReq, err error) {
	msg = &MessageBrowseReq{}
	return msg, nil
}

func (m *MessageBrowseReq) Length() int {
	length := 0
	length += LenMessageType
	length += LenChunkType
	length += LenMessageSize
	return length
}

func (m *MessageBrowseReq) Buffer() (*buffer.Buffer, error) {
	buf := buffer.NewBuffer(m.Length())
	if _, err := buf.Write([]byte{'M', 'S', 'G'}); err != nil {
		return nil, err
	}
	if _, err := buf.Write([]byte{'F'}); err != nil {
		return nil, err
	}
	return buf, nil
}
