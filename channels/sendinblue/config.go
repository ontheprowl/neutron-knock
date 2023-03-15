package channels

import (
	"context"
	"log"

	sendinblue "github.com/sendinblue/APIv3-go-library/v2/lib"
	"neutron.money/knock/types"
)

func (sib *SendInBlueProvider) Init() {

	cfg := sendinblue.NewConfiguration()
	//Configure API key authorization: api-key
	cfg.AddDefaultHeader("api-key", sib.ApiKey)
	//Configure API key authorization: partner-key
	cfg.AddDefaultHeader("partner-key", sib.PartnerKey)
	sib.client = sendinblue.NewAPIClient(cfg)
}

func (sib *SendInBlueProvider) SendMessage(content *types.Payload, jt types.JobType) {
	log.Println("Sending Email...")
	sibDataMap := content.Data.(map[string]interface{})
	templateID := int64(sibDataMap["templateID"].(float64))
	if jt == types.Email {
		sib.client.TransactionalEmailsApi.SendTransacEmail(context.Background(), sendinblue.SendSmtpEmail{
			Sender: &sendinblue.SendSmtpEmailSender{
				Name:  "Team Neutron",
				Email: "team@neutron.money",
			},
			To: []sendinblue.SendSmtpEmailTo{
				{
					Email: content.Contact,
				},
			},
			Params:     sibDataMap["params"].(map[string]interface{}),
			TemplateId: templateID,
		})
	} else {
		sib.client.TransactionalEmailsApi.SendTransacEmail(context.Background(), sendinblue.SendSmtpEmail{
			Sender: &sendinblue.SendSmtpEmailSender{
				Name:  "Team Neutron",
				Email: "team@neutron.money",
			},
			To: []sendinblue.SendSmtpEmailTo{
				{
					Email: content.Contact,
					Name:  content.ContactName,
				},
			},
			Params:     sibDataMap["params"].(map[string]interface{}),
			TemplateId: templateID,
		})
	}

}
