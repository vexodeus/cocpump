package coc

type Key struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Cidrranges  []string `json:"cidrRanges"`
	Key         string   `json:"key"`
}
