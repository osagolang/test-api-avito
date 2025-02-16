package models

type Shop struct {
	ID    int    `json:"id"`
	Type  string `json:"type"`
	Price int    `json:"price"`
}
