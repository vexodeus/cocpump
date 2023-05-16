package main

import (
	"encoding/json"

	resty "github.com/go-resty/resty/v2"
)

const (
	BaseUrl       = "https://developer.clashofclans.com"
	LoginEndpoint = "/api/login"
)

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Status                  Status `json:"status"`
	SessionExpiresInSeconds int    `json:"sessionExpiresInSeconds"`
	TemporaryAPIToken       string `json:"temporaryAPIToken"`
}
type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  any    `json:"detail"`
}

func login(email string, password string, client *resty.Client) (*LoginResponse, error) {
	//
	loginBody, err := json.Marshal(
		&LoginBody{
			Email:    email,
			Password: password,
		})
	if err != nil {
		return nil, err
	}
	//
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(loginBody).
		Post(BaseUrl + LoginEndpoint)
	if err != nil {
		return nil, err
	}
	//
	loginResponse := &LoginResponse{}
	err = json.Unmarshal(resp.Body(), loginResponse)
	if err != nil {
		return nil, err
	}
	//
	return loginResponse, nil
}
