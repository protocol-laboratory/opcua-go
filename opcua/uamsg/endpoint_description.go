package uamsg

type EndpointDescription struct {
	EndpointUrl         string
	Server              *ApplicationDescription
	ServerCertificate   []byte
	SecurityMode        MessageSecurityModeEnum
	SecurityPolicyUri   string
	UserIdentityTokens  []*UserTokenPolicy
	TransportProfileUri string
	SecurityLevel       byte
}
