package models

type Inventory struct {
	ID       int `json:"id"`
	UserID   int `json:"user_id"`
	ItemID   int `json:"item_id"`
	Quantity int `json:"quantity"`
}
