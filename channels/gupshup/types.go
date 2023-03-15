package channels

type GupshupProvider struct {
	ApiKey    string
	templates []template
	SrcName   string
	Source    string
}

type template struct {
	Id       string `json:"id"`
	Category string `json:"category"`
	Name     string `json:"elementName"`
	Status   string `json:"status"`
	Vertical string `json:"vertical"`
	Type     string `json:"templateType"`
}

type templateResponse struct {
	Status    string
	Templates []template
}

// "createdOn": 1677848088297,
// "data": "You have an upcoming payment due\nDear {{1}},\n\nHope this message finds you well. We are reaching out to kindly remind you that you have an upcoming payment due in {{5}} of amount {{2}} to {{3}}.\n\nWe value your business and appreciate your timely payments, which enable us to continue providing you with the quality products and services you deserve.\n\nIf you have any questions or concerns about your payment or billing, please do not hesitate to reach out to us. We are here to help and want to ensure that your experience with us is as smooth and stress-free as possible.\n\nThank you for your prompt attention to this matter, and we look forward to continuing our partnership with you.\n\nBest regards,\n{{3}}\nPowered By Neutron- Accounts Receivables On Autopilot  | [Visit Neutron,neutron.money]",
// "elementName": "payment_reminder",
// "id": "0d057ee5-75f8-456c-8ddf-a75f7a930958",
// "internalCategory": 0,
// "internalType": 0,
// "languageCode": "en",
// "languagePolicy": "deterministic",
// "meta": "{\"example\":\"Dear Customer,\\n\\nHope this message finds you well. We are reaching out to kindly remind you that you have an upcoming payment due in 2 days of amount 10,000 to Seller.\\n\\nWe value your business and appreciate your timely payments, which enable us to continue providing you with the quality products and services you deserve.\\n\\nIf you have any questions or concerns about your payment or billing, please do not hesitate to reach out to us. We are here to help and want to ensure that your experience with us is as smooth and stress-free as possible.\\n\\nThank you for your prompt attention to this matter, and we look forward to continuing our partnership with you.\\n\\nBest regards,\\nSeller\"}",
// "modifiedOn": 1677848088547,
// "namespace": "6f1b7388_ddbd_4c73_9475_e1818debc786",
// "priority": 1,
// "quality": "UNKNOWN",
// "reason": "component of type BODY is missing expected field(s) (example)",
// "retry": 0,
// "stage": "NONE",
// "status": "REJECTED",
// "templateType": "TEXT",
// "vertical": "Payment Reminder",
// "wabaId": "107743742257437"
