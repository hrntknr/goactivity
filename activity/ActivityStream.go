package activity

type ActivityStream struct {
	Type              string     `json:"type"`
	ID                string     `json:"id"`
	Name              string     `json:"name"`
	Following         string     `json:"following"`
	Followers         string     `json:"followers"`
	PreferredUsername string     `json:"preferredUsername"`
	Summary           string     `json:"summary"`
	Inbox             string     `json:"inbox"`
	Outbox            string     `json:"outbox"`
	URL               string     `json:"url"`
	Tag               []string   `json:"tag"`
	Attachment        []string   `json:"attachment"`
	Endpoints         *Endpoints `json:"endpoints"`
	Icon              *Icon      `json:"icon"`
}

func NewActivityStream(act *Activity) *ActivityStream {
	as := &ActivityStream{
		Endpoints: newEndpoints(),
		Icon:      newIcon(),
	}
	act.addContext(TypeActivityStreams)
	act.setActivityStream(as)
	return as
}
