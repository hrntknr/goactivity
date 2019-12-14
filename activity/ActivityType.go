package activity

type ActivityType int

var (
	TypeActivityStreams = registerActivityType(0, "https://www.w3.org/ns/activitystreams")
	TypeSecurity        = registerActivityType(1, "https://w3id.org/security/v1")
)

var contextMap = map[ActivityType]interface{}{}

func registerActivityType(num int, ctx interface{}) ActivityType {
	at := ActivityType(num)
	contextMap[at] = ctx
	return at
}
