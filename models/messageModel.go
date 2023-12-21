package models

type MessageQueue struct {
	Id      string                 `json:"id"`
	Pattern string                 `json:"pattern"`
	Data    map[string]interface{} `json:"data"`
}
