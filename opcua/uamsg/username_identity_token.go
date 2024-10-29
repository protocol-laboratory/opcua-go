package uamsg

type UserNameIdentityToken struct {
	PolicyId            string
	UserName            string
	Password            []byte
	EncryptionAlgorithm string
}
