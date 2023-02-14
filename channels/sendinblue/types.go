package channels

import (
	sendinblue "github.com/sendinblue/APIv3-go-library/v2/lib"
)

type SendInBlueProvider struct {
	ApiKey     string
	PartnerKey string
	client     *sendinblue.APIClient
}
