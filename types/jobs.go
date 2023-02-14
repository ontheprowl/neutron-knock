package types

type JobType int

type Payload struct {
	Message     string
	Data        map[string]interface{}
	Contact     string
	ContactName string
}

const (
	Email    = 0
	Whatsapp = 1
	Other    = 2
)

func (jt JobType) String() string {
	return []string{"Email", "Whatsapp", "Other"}[jt]
}

type Job struct {
	Id      string   `json:"id" xml:"id" form:"id"`
	JobType JobType  `json:"jobType" xml:"jobType" form:"jobType"`
	Data    *Payload `json:"jobType" xml:"jobType" form:"jobType"`
}
