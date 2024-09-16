package opcua

import "github.com/shoothzj/gox/buffer"

type MessageHello struct {
	Version           uint32
	ReceiveBufferSize uint32
	SendBufferSize    uint32
	MaxMessageSize    uint32
	MaxChunkCount     uint32
	EndpointUrl       string
}

func DecodeMessageHello(buf *buffer.Buffer) (msg *MessageHello, err error) {
	msg = &MessageHello{}
	msg.Version, err = buf.ReadUInt32Le()
	if err != nil {
		return nil, err
	}
	msg.ReceiveBufferSize, err = buf.ReadUInt32Le()
	if err != nil {
		return nil, err
	}
	msg.SendBufferSize, err = buf.ReadUInt32Le()
	if err != nil {
		return nil, err
	}
	msg.MaxMessageSize, err = buf.ReadUInt32Le()
	if err != nil {
		return nil, err
	}
	msg.MaxChunkCount, err = buf.ReadUInt32Le()
	if err != nil {
		return nil, err
	}
	msg.EndpointUrl, err = buf.ReadStringLe()
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (m *MessageHello) Length() int {
	length := 0
	length += LenMessageType
	length += LenChunkType
	length += LenMessageSize
	length += LenVersion
	length += LenReceiveBufferSize
	length += LenSendBufferSize
	length += LenMaxMessageSize
	length += LenMaxChunkCount
	length += StrLen(m.EndpointUrl)
	return length
}

func (m *MessageHello) Buffer() (*buffer.Buffer, error) {
	buf := buffer.NewBuffer(m.Length())
	if _, err := buf.Write([]byte{'H', 'E', 'L'}); err != nil {
		return nil, err
	}
	if _, err := buf.Write([]byte{'F'}); err != nil {
		return nil, err
	}
	if err := buf.PutUInt32Le(uint32(m.Length())); err != nil {
		return nil, err
	}
	if err := buf.PutUInt32Le(m.Version); err != nil {
		return nil, err
	}
	if err := buf.PutUInt32Le(m.ReceiveBufferSize); err != nil {
		return nil, err
	}
	if err := buf.PutUInt32Le(m.SendBufferSize); err != nil {
		return nil, err
	}
	if err := buf.PutUInt32Le(m.MaxMessageSize); err != nil {
		return nil, err
	}
	if err := buf.PutUInt32Le(m.MaxChunkCount); err != nil {
		return nil, err
	}
	if err := buf.PutStringLe(m.EndpointUrl); err != nil {
		return nil, err
	}
	return buf, nil
}
