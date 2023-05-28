package api

import "cocpump/pkg/coc"

type LoginResponse struct {
	Status                  Status `json:"status"`
	SessionExpiresInSeconds int    `json:"sessionExpiresInSeconds"`
	TemporaryAPIToken       string `json:"temporaryAPIToken"`
}
type KeyResponse struct {
	Status                  Status  `json:"status"`
	SessionExpiresInSeconds int     `json:"sessionExpiresInSeconds"`
	Key                     coc.Key `json:"key"`
}
type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  any    `json:"detail"`
}
