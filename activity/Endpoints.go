package activity

type Endpoints struct {
	SharedInbox string `json:"sharedInbox"`
}

func newEndpoints() *Endpoints {
	return &Endpoints{}
}
