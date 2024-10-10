package uamsg

type HelloMessageExtras struct {
	ProtocolVersion   uint32
	ReceiveBufferSize uint32
	SendBufferSize    uint32
	MaxMessageSize    uint32
	MaxChunkCount     uint32
	EndpointUrl       string
}

type AcknowledgeMessageExtras struct {
	ProtocolVersion   uint32
	ReceiveBufferSize uint32
	SendBufferSize    uint32
	MaxMessageSize    uint32
	MaxChunkCount     uint32
}

type ReverseHelloMessageExtras struct {
	ServerUri   string
	EndpointUrl string
}
