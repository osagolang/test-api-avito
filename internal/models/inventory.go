package models

type Inventory struct {
	ID       int    `json:"-"`
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}
