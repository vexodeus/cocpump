package api

import (
	"encoding/json"

	resty "github.com/go-resty/resty/v2"
)

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
