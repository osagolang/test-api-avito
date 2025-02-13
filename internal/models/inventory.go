package models

type Inventory struct {
	ID       int `json:"id"`
	User_id  int `json:"user_id"`
	Item_id  int `json:"item_id"`
	Quantity int `json:"quantity"`
}
