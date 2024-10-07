package uamsg

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
	// Message的序列化需要先计算MessageService的大小，然后计算Header的值，最后完成整个序列化
	// 解码则需要根据Header字段分析需要跳过什么内容
	MessageBody interface{}

	// 这个基本可以无视了，这个第一把先不搞
	MessageFooter interface{}
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

// Error消息的附加字段，任何服务端响应报错都应该考虑在MessageBody里头添加这个
type ErrorMessageExtras struct {
	Error  uint32
	Reason string
}

type RequestHeader struct {
	AuthenticationToken *NodeId // session创建后，需要服务端生成一个字符串，使用NodeId承载

	Timestamp         uint64 // 编码为uint64 8个字节
	RequestHandle     IntegerId
	ReturnDiagnostics uint32
	AuditEntryId      string
	TimeoutHint       uint32
	AdditionalHeader  *ExtensionObject
}

type ResponseHeader struct {
	Timestamp          uint64 // 编码为uint64 8个字节
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
	Namespace    uint16             // 这个字段在不同结构中长度不一样
	Identifier   interface{}
}

type ExpandedNodeId struct {
	NodeId
	NamespaceUri string
	ServerIndex  uint32
}

type SessionAuthenticationToken interface{}

type ExtensionObject struct {
	TypeId   NodeId
	Encoding byte
	Length   int32
	Body     string // 可选字段
}

type GenericBody struct {
	TypeId  *ExpandedNodeId
	Service interface{}
}
