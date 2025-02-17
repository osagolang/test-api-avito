package repositories

import (
	"database/sql"
	"test-api-avito/internal/models"
)

type ItemRepo struct {
	db *sql.DB
}

func NewItemRepo(db *sql.DB) *ItemRepo {
	return &ItemRepo{db: db}
}

// Получить товар по названию
func (r *ItemRepo) GetItem(name string) (*models.Item, error) {
	var item models.Item
	err := r.db.QueryRow("SELECT type, price FROM items WHERE name = $1", name).Scan(&item.Type, &item.Price)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
