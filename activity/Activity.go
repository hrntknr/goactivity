package activity

type Activity struct {
	Context []interface{} `json:"@context"`
	*ActivityStream
	*Security
}

func NewActivity() *Activity {
	return &Activity{}
}

func (act *Activity) addContext(at ActivityType) {
	act.Context = append(act.Context, contextMap[at])
}

func (act *Activity) setActivityStream(as *ActivityStream) {
	act.ActivityStream = as
}

func (act *Activity) setSecurity(sec *Security) {
	act.Security = sec
}
