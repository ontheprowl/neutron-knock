package channels

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"neutron.money/knock/types"
)

func (gs *GupshupProvider) Init() {

	templates, err := getAllTemplates(gs.SrcName, gs.ApiKey)
	if err != nil {
		log.Println("Could not initialize Gupshup templates")
	}

	gs.templates = templates

}

func getAllTemplates(appName string, apikey string) ([]template, error) {
	url := `https://api.gupshup.io/sm/api/v1/template/list/` + appName
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("apikey", apikey)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var templateResponse templateResponse

	err = json.Unmarshal(body, &templateResponse)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return templateResponse.Templates, nil

}

func (gs *GupshupProvider) SendMessage(content *types.Payload, jt types.JobType) {
	url := "http://api.gupshup.io/sm/api/v1/template/msg"
	method := "POST"

	paramsArray := content.Data.([]interface{})

	paramsBytes, err := json.Marshal(paramsArray)

	if err != nil {
		log.Println("Error encountered while marshalling Gupshup Params Array")
	}

	payloadString := fmt.Sprintf(`source=919175121966&destination=%s&template={"id":"%s","params":%s}&channel=whatsapp`, content.Contact, content.Message, string(paramsBytes))

	payload := strings.NewReader(payloadString)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Close = true
	req.Header.Add("apikey", gs.ApiKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(body))
}
