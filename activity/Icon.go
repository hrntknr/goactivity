package activity

type Icon struct {
	Type      string `json:"type"`
	MediaType string `json:"mediaType"`
	URL       string `json:"url"`
}

func newIcon() *Icon {
	return &Icon{}
}
