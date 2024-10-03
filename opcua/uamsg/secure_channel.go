package uamsg

type OpenSecureChannelServiceRequest struct {
	Header *RequestHeader
	ClientProtocolVersion uint32
	RequestType           SecurityTokenRequestType
	SecurityMode          MessageSecurityModeEnum
	ClientNonce           []byte
	RequestedLifetime     uint32
}

type OpenSecureChannelServiceResponse struct {
	Header *ResponseHeader
	ServerProtocolVersion uint32
	SecurityToken         ChannelSecurityToken // todo 没找到这个定义，需要去查一查
	SecureChannelId       uint32
	TokenId               uint32
	CreatedAt             uint64
	RevisedLifetime       uint32
	ServerNonce           []byte //
}

type DiagnosticInfo struct {
	NamespaceUri        int32
	SymbolicId          int32
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
	SecureChannelId BaseDataType
}

type CloseSecureChannelResponse struct {
	Header *ResponseHeader
}
