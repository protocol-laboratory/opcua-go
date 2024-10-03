package uamsg

import (
	"time"
)

type CreateSessionRequest struct {
	Header *RequestHeader
	ClientDescription       *ApplicationDescription
	ServerUri               string
	EndpointUrl             string
	SessionName             string
	ClientNonce             []byte
	ClientCertificate       []byte        // 客户端证书，自己配吧，就按照golang的来解析
	RequestedSessionTimeout time.Duration // 编码的时候转为uint64
	MaxResponseMessageSize  uint32
}

type ApplicationDescription struct {
	ApplicationUri      string
	ProductUri          string
	ApplicationName     LocalizedText
	ApplicationType     ApplicationTypeEnum
	GatewayServerUri    string
	DiscoveryProfileUri string
	DiscoveryUrls       []string
}

type LocalizedText struct {
	EncodingMask byte
	Locale       string
	Text         string
}

type CreateSessionResponse struct {
	Header *ResponseHeader
	SessionId                  *NodeId
	AuthenticationToken        SessionAuthenticationToken
	RevisedSessionTimeout      uint64 // 编码的时候转为uint64
	ServerNonce                []byte
	ServerCertificate          []byte // https://reference.opcfoundation.org/Core/Part4/v105/docs/7.3#_Ref182127421 结构不好定义
	ServerEndpoints            []EndpointDescription
	ServerSoftwareCertificates []SignedSoftwareCertificate
	ServerSignature            SignatureData
	MaxRequestMessageSize      uint32
}

type EndpointDescription struct {
	EndpointUrl         string
	Server              ApplicationDescription
	ServerCertificate   []byte
	SecurityMode        MessageSecurityModeEnum
	SecurityPolicyUri   string
	UserIdentityTokens  []UserTokenPolicy
	TransportProfileUri string
	SecurityLevel       byte
}

type UserTokenPolicy struct {
	PolicyId          string
	TokenType         UserTokenTypeEnum
	IssuedTokenType   string
	IssuerEndpointUrl string
	SecurityPolicyUri string
}

type SignedSoftwareCertificate struct {
	CertificateData []byte
	Signature       []byte
}

type SignatureData struct {
	Algorithm string
	Signature []byte
}

type ActivateSessionRequest struct {
	Header *RequestHeader
	ClientSignature            SignatureData
	ClientSoftwareCertificates []SignedSoftwareCertificate
	LocaleIds                  []LocalId
	UserIdentityToken          *ExtensibleParameter
	UserTokenSignature         SignatureData
}

type ExtensibleParameter struct {
	ParameterTypeId NodeId
	ParameterData   interface{}
}

type ActivateSessionResponse struct {
	ServerNonce     []byte
	Results         []StatusCode
	DiagnosticInfos []DiagnosticInfo
}

type CloseSessionRequest struct {
	Header *RequestHeader
	DeleteSubscriptions bool
}

type CloseSessionResponse struct {
	Header *ResponseHeader
}
