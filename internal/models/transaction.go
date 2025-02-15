package models

import "time"

type Transaction struct {
	ID         int       `json:"id"`
	FromUserID int       `json:"from_user"`
	ToUserID   int       `json:"to_user"`
	Amount     int       `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
}
