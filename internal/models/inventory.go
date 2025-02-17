package models

type Inventory struct {
	ID       int `json:"id"`
	Type     int `json:"type"`
	Quantity int `json:"quantity"`
}
