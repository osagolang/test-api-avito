package models

type Wallet struct {
	UserId  int `json:"user_id"`
	Balance int `json:"balance"`
}
