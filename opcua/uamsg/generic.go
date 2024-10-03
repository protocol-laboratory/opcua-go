package uamsg

import "time"

type (
	SecurityTokenRequestType uint32
	MessageSecurityModeEnum  uint32

	BaseDataType           interface{}
	IntegerId              uint32
	StatusCode             uint32
	ApplicationTypeEnum    uint32
	UserTokenTypeEnum      uint32
	LocalId                string
	TimestampsToReturnEnum uint32
	NumericRange           string
)

type Message struct {
	MessageHeader
	SecurityHeader interface{}
	SequenceHeader
	// Message的序列化需要先计算MessageService的大小，然后计算Header的值，最后完成整个序列化
	// 解码则需要根据Header字段分析需要跳过什么内容
	MessageService interface{}

	// 这个基本可以无视了，这个第一把先不搞
	MessageFooter interface{}
}

type MessageHeader struct {
	MessageType [3]byte
	ChunkType   byte
	MessageSize uint32
	SecureChannelId uint32 // handshake阶段，secure channel未建立，无需编码该字段	
}

type AsymmetricSecurityHeader struct {
	SecurityPolicyUriLength             int32
	SecurityPolicyUri                   []byte
	SenderCertificateLength             int32
	SenderCertificate                   []byte
	ReceiverCertificateThumbprintLength int32
	ReceiverCertificateThumbprint       []byte
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

	Timestamp         time.Time // 编码为uint64 8个字节
	RequestHandle     IntegerId
	ReturnDiagnostics uint32
	AuditEntryId      string
	TimeoutHint       uint32
	AdditionalHeader  *AdditionalParametersType
}

type ResponseHeader struct {
	Timestamp          time.Time // 编码为uint64 8个字节
	RequestHandle      IntegerId
	ServiceResult      StatusCode
	ServiceDiagnostics *DiagnosticInfo
	StringTable        []string
	AdditionalHeader   *AdditionalParametersType
}

type QualifiedName struct {
	NamespaceIndex uint16
	Name           string
}

type NodeId struct {
	EncodingType   uint16 // indicate nodeId format
	Namespace      byte
	IdentifierType uint32
	Identifier     interface{}
}

type AdditionalParametersType struct {
	Parameters map[QualifiedName]BaseDataType
}

type SessionAuthenticationToken interface{}
