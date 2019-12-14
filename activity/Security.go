package activity

type Security struct {
	PublicKey *PublicKey `json:"publicKey"`
}

func NewSecurity(act *Activity) *Security {
	sec := &Security{
		PublicKey: newPublicKey(),
	}
	act.addContext(TypeSecurity)
	act.setSecurity(sec)
	return sec
}
