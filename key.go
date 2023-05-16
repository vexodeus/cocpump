package main

import (
	"encoding/json"

	resty "github.com/go-resty/resty/v2"
)

const KeyListEndpoint = "/api/apikey/list"

type KeyResponse struct {
	Status                  Status `json:"status"`
	SessionExpiresInSeconds int    `json:"sessionExpiresInSeconds"`
	Key                     Key    `json:"key"`
}

type Key struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Cidrranges  []string `json:"cidrRanges"`
	Key         string   `json:"key"`
}

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
