package model

type ValidateIPRequest struct {
	Ip string `json:"ip"`
}

type ValidateIPResponse struct {
	Status bool `json:"status"`
}
