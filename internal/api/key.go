package api

import (
	"encoding/json"

	resty "github.com/go-resty/resty/v2"
)

func getKeys(email string, password string, client *resty.Client) (*KeyResponse, error) {
	//
	loginResponse, err := login(email, password, client)
	if err != nil {
		return nil, err
	}
	token := loginResponse.TemporaryAPIToken
	//
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(token).
		Post(BaseUrl + KeyListEndpoint)
	if err != nil {
		return nil, err
	}
	//
	keyResponse := &KeyResponse{}
	err = json.Unmarshal(resp.Body(), keyResponse)
	if err != nil {
		return nil, err
	}
	//
	return keyResponse, nil
}
