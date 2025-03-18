package models

type Transaction struct {
	ID     int    `json:"-"`
	User   string `json:"user"`
	Amount int    `json:"amount"`
}

type CoinHistory struct {
	Received []Transaction `json:"received"`
	Sent     []Transaction `json:"sent"`
}
