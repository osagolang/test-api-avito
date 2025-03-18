package models

type Item struct {
	ID    int    `json:"id"`
	Type  string `json:"type"`
	Price int    `json:"price"`
}
