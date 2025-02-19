package repositories

import (
	"database/sql"
	"errors"
)

type BuyRepo struct {
	db       *sql.DB
	itemRepo *ItemRepo
}

func NewBuyRepo(db *sql.DB, itemRepo *ItemRepo) *BuyRepo {
	return &BuyRepo{db: db, itemRepo: itemRepo}
}

// Покупка предмета
func (r *BuyRepo) BuyItem(userID int, itemName string) error {

	// Проверяем товар
	item, err := r.itemRepo.GetItem(itemName)
	if err != nil {
		return errors.New("Товар не найден.")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Проверяем баланс пользователя
	var userCoins int
	err = tx.QueryRow("SELECT coins FROM users WHERE id = $1 FOR UPDATE", userID).Scan(&userCoins)
	if err != nil {
		return err
	}
	if userCoins < item.Price {
		return errors.New("Недостаточно монет")
	}

	// Вычитаем монеты
	_, err = tx.Exec("UPDATE users SET coins = coins - $1 WHERE id = $2", item.Price, userID)
	if err != nil {
		return err
	}

	// Добавляем товар в инвентарь
	_, err = tx.Exec("INSERT INTO inventory (user_id, type, quantity) VALUES ($1, $2, 1) ON CONFLICT (user_id, type) DO UPDATE SET quantity = inventory.quantity + 1", userID, item.Type)
	if err != nil {
		return err
	}

	// Фиксируем транзакцию
	return tx.Commit()
}
