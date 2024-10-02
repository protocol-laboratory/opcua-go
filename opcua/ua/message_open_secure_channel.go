package ua

import (
	"github.com/shoothzj/gox/buffer"
)

type MessageOpenSecureChannel struct {
	SecureChannelId               uint32
	SecurityPolicyUri             string
	SenderCertificate             []byte
	ReceiverCertificateThumbprint string
	SequenceNumber                uint32
	RequestId                     uint32
}

func DecodeMessageOpenSecureChannel(buf *buffer.Buffer) (msg *MessageOpenSecureChannel, err error) {
	msg = &MessageOpenSecureChannel{}
	return msg, nil
}

func (m *MessageOpenSecureChannel) Length() int {
	length := 0
	length += LenMessageType
	length += LenChunkType
	length += LenMessageSize
	length += StrLen(m.SecurityPolicyUri)
	// todo senderCertificate
	length += 4
	// todo receiverCertificateThumbprint
	length += 4
	length += LenSequenceNumber
	length += LenRequestId
	return length
}

func (m *MessageOpenSecureChannel) Buffer() (*buffer.Buffer, error) {
	buf := buffer.NewBuffer(m.Length())
	if _, err := buf.Write([]byte{'O', 'P', 'N'}); err != nil {
		return nil, err
	}
	if _, err := buf.Write([]byte{'F'}); err != nil {
		return nil, err
	}
	if err := buf.PutUInt32Le(uint32(m.Length())); err != nil {
		return nil, err
	}
	return buf, nil
}
