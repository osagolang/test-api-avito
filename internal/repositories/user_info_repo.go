package repositories

import (
	"database/sql"
	"errors"
	"test-api-avito/internal/models"
)

type UserInfoRepo struct {
	db *sql.DB
}

func NewUserInfoRepo(db *sql.DB) *UserInfoRepo {
	return &UserInfoRepo{db: db}
}

// Открытие новой транзакции
func (r *UserInfoRepo) BeginTransaction() (*sql.Tx, error) {
	return r.db.Begin()
}

// UserCoins возвращает количество монет пользователя
func (r *UserInfoRepo) UserCoins(userID int) (int, error) {
	var coins int
	err := r.db.QueryRow("SELECT coins FROM users WHERE id = $1", userID).Scan(&coins)
	return coins, err
}

// UserInventory возвращает инвентарь пользователя
func (r *UserInfoRepo) UserInventory(userID int) ([]models.Inventory, error) {
	rows, err := r.db.Query("SELECT item, quantity FROM inventory WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventory []models.Inventory
	for rows.Next() {
		var item models.Inventory
		if err := rows.Scan(&item.Type, &item.Quantity); err != nil {
			return nil, err
		}
		inventory = append(inventory, item)
	}
	return inventory, nil
}

// UserTransactionHistory возвращает историю переводов монет
func (r *UserInfoRepo) UserTransactionHistory(userID int) (models.CoinHistory, error) {
	var history models.CoinHistory

	receivedRows, err := r.db.Query("SELECT from_user, amount FROM transactions WHERE to_user = $1", userID)
	if err != nil {
		return history, err
	}
	defer receivedRows.Close()

	for receivedRows.Next() {
		var tx models.Transaction
		if err := receivedRows.Scan(&tx.User, &tx.Amount); err != nil {
			return history, err
		}
		history.Received = append(history.Received, tx)
	}

	sentRows, err := r.db.Query("SELECT to_user, amount FROM transactions WHERE from_user = $1", userID)
	if err != nil {
		return history, err
	}
	defer sentRows.Close()

	for sentRows.Next() {
		var tx models.Transaction
		if err := sentRows.Scan(&tx.User, &tx.Amount); err != nil {
			return history, err
		}
		history.Sent = append(history.Sent, tx)
	}

	return history, nil
}

// Перевод монет (с транзакцией)
func (r *UserInfoRepo) TransferCoins(fromUserID, toUserID, amount int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Блокируем баланс отправителя
	var senderBalance int
	err = tx.QueryRow("SELECT coins FROM users WHERE id = $1 FOR UPDATE", fromUserID).Scan(&senderBalance)
	if err != nil {
		return err
	}

	if senderBalance < amount {
		return errors.New("недостаточно монет")
	}

	// Вычитаем монеты у отправителя
	_, err = tx.Exec("UPDATE users SET coins = coins - $1 WHERE id = $2", amount, fromUserID)
	if err != nil {
		return err
	}

	// Добавляем монеты получателю
	_, err = tx.Exec("UPDATE users SET coins = coins + $1 WHERE id = $2", amount, toUserID)
	if err != nil {
		return err
	}

	// Записываем перевод в историю
	_, err = tx.Exec("INSERT INTO transactions (from_user, to_user, amount) VALUES ($1, $2, $3)", fromUserID, toUserID, amount)
	if err != nil {
		return err
	}

	return tx.Commit()
}
