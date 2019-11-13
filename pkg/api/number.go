package api

type Number struct {
	Number int `json:"number"`
	Classification string `json:"classification,omitempty"`
}
