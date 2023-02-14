package channels

import (
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

	log.Println("Hello World")
	log.Println(content)
	log.Println(jt)
	// if jt == types.Email {
	// 	sib.client.TransactionalEmailsApi.SendTransacEmail(context.Background(), sendinblue.SendSmtpEmail{
	// 		Sender: &sendinblue.SendSmtpEmailSender{
	// 			Name:  "Team Neutron",
	// 			Email: "team@neutron.money",
	// 		},
	// 		To: []sendinblue.SendSmtpEmailTo{
	// 			{
	// 				Email: content.Contact,
	// 				Name:  content.ContactName,
	// 			},
	// 		},
	// 		Params:     content.Data["params"].(map[string]interface{}),
	// 		TemplateId: content.Data["templateID"].(int64),
	// 	})
	// } else {
	// 	sib.client.TransactionalEmailsApi.SendTransacEmail(context.Background(), sendinblue.SendSmtpEmail{
	// 		Sender: &sendinblue.SendSmtpEmailSender{
	// 			Name:  "Team Neutron",
	// 			Email: "team@neutron.money",
	// 		},
	// 		To: []sendinblue.SendSmtpEmailTo{
	// 			{
	// 				Email: content.Contact,
	// 				Name:  content.ContactName,
	// 			},
	// 		},
	// 		Params:     content.Data["params"].(map[string]interface{}),
	// 		TemplateId: content.Data["templateID"].(int64),
	// 	})
	// }

}
