package uamsg

type MessageTypeEnum [3]byte

var (
	HelloMessageType              MessageTypeEnum = [3]byte{'H', 'E', 'L'}
	AcknowledgeMessageType        MessageTypeEnum = [3]byte{'A', 'C', 'K'}
	OpenSecureChannelMessageType  MessageTypeEnum = [3]byte{'O', 'P', 'N'}
	MsgMessageType                MessageTypeEnum = [3]byte{'M', 'S', 'G'}
	CloseSecureChannelMessageType MessageTypeEnum = [3]byte{'C', 'L', 'O'}
)

type ChunkTypeEnum byte

const (
	AbortChunkType        ChunkTypeEnum = 'A'
	IntermediateChunkType ChunkTypeEnum = 'C'
	FinalChunkType        ChunkTypeEnum = 'F'
)

type NodeIdEncodingType byte

const (
	TwoByte          NodeIdEncodingType = 0x00
	FourByte         NodeIdEncodingType = 0x01
	Numeric          NodeIdEncodingType = 0x02
	String           NodeIdEncodingType = 0x03
	GuidType         NodeIdEncodingType = 0x04
	ByteString       NodeIdEncodingType = 0x05
	NamespaceUriFlag NodeIdEncodingType = 0x80
	ServerIndexFlag  NodeIdEncodingType = 0x40
)

type Guid struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 uint64
}

type DiagnosticInfoMaskEnum = byte

const (
	SymbolicIdFlag          DiagnosticInfoMaskEnum = 0x01
	NamespaceFlag           DiagnosticInfoMaskEnum = 0x02
	LocalizedTextFlag       DiagnosticInfoMaskEnum = 0x04
	LocaleFlag              DiagnosticInfoMaskEnum = 0x08
	AdditionalInfoFlag      DiagnosticInfoMaskEnum = 0x10
	InnerStatusCodeFlag     DiagnosticInfoMaskEnum = 0x20
	InnerDiagnosticInfoFlag DiagnosticInfoMaskEnum = 0x40
)

type ServiceTypeEnum uint16

const (
	OpenSecureChannelServiceRequestType  ServiceTypeEnum = 446
	OpenSecureChannelServiceResponseType ServiceTypeEnum = 449
	CreateSessionRequestType             ServiceTypeEnum = 461
	CreateSessionResponseType            ServiceTypeEnum = 464
)
