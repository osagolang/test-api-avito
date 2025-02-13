package models

import "time"

type Transaction struct {
	ID        int       `json:"id"`
	FromUser  int       `json:"from_user"`
	ToUser    int       `json:"to_user"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
