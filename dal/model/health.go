package model

type Health struct {
	Server   string `json:"server"`
	Metrics  string `json:"metrics"`
	Database string `json:"database"`
}
