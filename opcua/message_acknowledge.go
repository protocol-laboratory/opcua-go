package opcua

import "github.com/shoothzj/gox/buffer"

type MessageAcknowledge struct {
	Version           uint32
	ReceiveBufferSize uint32
	SendBufferSize    uint32
	MaxMessageSize    uint32
	MaxChunkCount     uint32
}

func DecodeMessageAcknowledge(buf *buffer.Buffer) (msg *MessageAcknowledge, err error) {
	msg = &MessageAcknowledge{}
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
	return msg, nil
}

func (m *MessageAcknowledge) Length() int {
	length := 0
	length += LenMessageType
	length += LenChunkType
	length += LenMessageSize
	length += LenVersion
	length += LenReceiveBufferSize
	length += LenSendBufferSize
	length += LenMaxMessageSize
	length += LenMaxChunkCount
	return length
}

func (m *MessageAcknowledge) Buffer() (*buffer.Buffer, error) {
	buf := buffer.NewBuffer(m.Length())
	if _, err := buf.Write([]byte{'A', 'C', 'K'}); err != nil {
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
	return buf, nil
}
