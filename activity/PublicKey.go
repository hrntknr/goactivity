package activity

type PublicKey struct {
	ID           string `json:"id"`
	Owner        string `json:"owner"`
	PublicKeyPem string `json:"publicKeyPem"`
}

func newPublicKey() *PublicKey {
	return &PublicKey{}
}
