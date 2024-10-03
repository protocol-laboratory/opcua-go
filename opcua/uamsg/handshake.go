package uamsg

// Hello消息的附加字段
type HelloMessageExtras struct {
	ProtocolVersion   uint32
	ReceiveBufferSize uint32
	SendBufferSize    uint32
	MaxMessageSize    uint32
	MaxChunkCount     uint32
	EndpointUrl       string
}

// Acknowledge消息的附加字段
type AcknowledgeMessageExtras struct {
	ProtocolVersion   uint32
	ReceiveBufferSize uint32
	SendBufferSize    uint32
	MaxMessageSize    uint32
	MaxChunkCount     uint32
}

// ReverseHello消息的附加字段，服务端反向连接客户端发出的第一个消息
type ReverseHelloMessageExtras struct {
	ServerUri   string
	EndpointUrl string
}
