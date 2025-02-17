package models

type Transaction struct {
	ID     int `json:"id"`
	User   int `json:"user"`
	Amount int `json:"amount"`
}

type CoinHistory struct {
	Received []Transaction `json:"received"`
	Sent     []Transaction `json:"sent"`
}
