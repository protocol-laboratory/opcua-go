package uamsg

const (
	MessageSecurityModeInvalid        MessageSecurityModeEnum = 0
	MessageSecurityModeNone           MessageSecurityModeEnum = 1
	MessageSecurityModeSign           MessageSecurityModeEnum = 2
	MessageSecurityModeSignAndEncrypt MessageSecurityModeEnum = 3
)

type (
	SecurityTokenRequestType uint32
	MessageSecurityModeEnum  uint32

	BaseDataType        interface{}
	IntegerId           uint32
	ApplicationTypeEnum uint32
	UserTokenTypeEnum   uint32

	TimestampsToReturnEnum uint32
	Duration               float64

	StatusCode   = uint32
	NumericRange = string
	LocalId      = string
)

type Message struct {
	*MessageHeader
	SecurityHeader interface{}
	*SequenceHeader
	MessageBody   interface{}
	MessageFooter interface{} // ignore
}

type MessageHeader struct {
	MessageType     MessageTypeEnum
	ChunkType       ChunkTypeEnum
	MessageSize     uint32
	SecureChannelId *uint32 `enc:"omitempty"` // handshake阶段，secure channel未建立，无需编码该字段
}

type AsymmetricSecurityHeader struct {
	SecurityPolicyUri             []byte
	SenderCertificate             []byte
	ReceiverCertificateThumbprint []byte
}

type SymmetricSecurityHeader struct {
	TokenId uint32
}

type SequenceHeader struct {
	SequenceNumber uint32
	RequestId      uint32
}

// ErrorMessageExtras Any server-side error responses are returned in this form
type ErrorMessageExtras struct {
	Error  uint32
	Reason string
}

type RequestHeader struct {
	AuthenticationToken *NodeId

	Timestamp         uint64
	RequestHandle     IntegerId
	ReturnDiagnostics uint32
	AuditEntryId      string
	TimeoutHint       uint32
	AdditionalHeader  *ExtensionObject
}

type ResponseHeader struct {
	Timestamp          uint64
	RequestHandle      IntegerId
	ServiceResult      StatusCode
	ServiceDiagnostics *DiagnosticInfo
	StringTable        []string
	AdditionalHeader   *ExtensionObject
}

type QualifiedName struct {
	NamespaceIndex uint16
	Name           string
}

type NodeId struct {
	EncodingType NodeIdEncodingType // indicate nodeId format
	Namespace    uint16
	Identifier   interface{}
}

type ExpandedNodeId struct {
	*NodeId
	NamespaceUri string
	ServerIndex  uint32
}

type SessionAuthenticationToken interface{}

type ExtensionObject struct {
	TypeId   *NodeId
	Encoding byte
	Body     string // optional, when Encoding=0x00, Body shouldn`t be passed
}

type GenericBody struct {
	TypeId  *ExpandedNodeId
	Service interface{}
}
