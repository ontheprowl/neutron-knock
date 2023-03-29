package channels

import (
	"errors"

	gupshup "neutron.money/knock/channels/gupshup"
	sendinblue "neutron.money/knock/channels/sendinblue"
	"neutron.money/knock/types"
)

var gupshupChannel, sibChannel types.ChannelProvider

// Initialize default channels for Neutron Knock. Current channels - SendInBlue & Gupshup
func InitDefaultChannels() {

	sibChannel = &sendinblue.SendInBlueProvider{
		ApiKey:     "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",
		PartnerKey: "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",
	}

	gupshupChannel = &gupshup.GupshupProvider{
		ApiKey:  "tisglowit8h33q2h33t4xtal8esza6z8",
		SrcName: "dQBruFMlIoDEWsL36AuspLCO"}

	sibChannel.Init()
	gupshupChannel.Init()
}

// Retrieves the appropriate communication channel for the provided jobtype
func GetChannel(jt types.JobType) (types.ChannelProvider, error) {
	switch jt {
	case 0:
		return sibChannel, nil
	case 1:
		return gupshupChannel, nil
	default:
		break
	}
	return nil, errors.New("no communication channel available for the specified jobtype")
}
