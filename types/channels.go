package types

type ChannelProvider interface {

	// This method initializes the communication channel
	Init()

	// This method sends a message through the communication channel
	SendMessage(content *Payload, jt JobType)
}
