package opcua

import "github.com/shoothzj/gox/buffer"

type MessageBrowseResp struct {
}

func DecodeMessageBrowseResp(buf *buffer.Buffer) (msg *MessageBrowseResp, err error) {
	msg = &MessageBrowseResp{}
	return msg, nil
}

func (m *MessageBrowseResp) Length() int {
	length := 0
	length += LenMessageType
	length += LenChunkType
	length += LenMessageSize
	return length
}

func (m *MessageBrowseResp) Buffer() (*buffer.Buffer, error) {
	buf := buffer.NewBuffer(m.Length())
	if _, err := buf.Write([]byte{'M', 'S', 'G'}); err != nil {
		return nil, err
	}
	if _, err := buf.Write([]byte{'F'}); err != nil {
		return nil, err
	}
	return buf, nil
}
