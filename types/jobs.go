package types

type JobType int

type Payload struct {
	Message     string      `json:"message" xml:"message" form:"message"`
	Data        interface{} `json:"data" xml:"data" form:"data"`
	Contact     string      `json:"contact" xml:"contact" form:"contact"`
	ContactName string      `json:"contactName" xml:"contactName" form:"contactName"`
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
	Id       string   `json:"id" xml:"id" form:"id"`
	JobType  JobType  `json:"jobType" xml:"jobType" form:"jobType"`
	Data     *Payload `json:"data" xml:"data" form:"data"`
	Schedule string   `json:"schedule" xml:"schedule" form:"schedule"`
}
