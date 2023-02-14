package types

type ChannelProvider interface {
	Init()
	SendMessage(content *Payload, jt JobType)
}
