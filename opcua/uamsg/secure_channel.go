package uamsg

type OpenSecureChannelServiceRequest struct {
	Header                *RequestHeader
	ClientProtocolVersion uint32
	RequestType           SecurityTokenRequestType
	SecurityMode          MessageSecurityModeEnum
	ClientNonce           []byte
	RequestedLifetime     uint32
}

type OpenSecureChannelServiceResponse struct {
	Header                *ResponseHeader
	ServerProtocolVersion uint32
	SecurityToken         *ChannelSecurityToken
	ServerNonce           []byte
}

type DiagnosticInfo struct {
	EncodingMask        DiagnosticInfoMaskEnum
	SymbolicId          int32
	NamespaceUri        int32
	Locale              int32
	LocalizedText       int32
	AdditionalInfo      string
	InnerStatusCode     StatusCode
	InnerDiagnosticInfo *DiagnosticInfo
}

type ChannelSecurityToken struct {
	ChannelID       uint32
	TokenID         uint32
	CreatedAt       uint64
	RevisedLifetime uint32
}

type CloseSecureChannelRequest struct {
	Header *RequestHeader
	// SecureChannelId *uint32 `enc:"omitempty"` // 协议规范定义了这个字段,但是别家实现不涉及
}

type CloseSecureChannelResponse struct {
	Header *ResponseHeader
}
