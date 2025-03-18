package repositories

import (
	"database/sql"
	"test-api-avito/internal/models"
)

type UserInfoRepo struct {
	db *sql.DB
}

func NewUserInfoRepo(db *sql.DB) *UserInfoRepo {
	return &UserInfoRepo{db: db}
}

// UserCoins возвращает количество монет пользователя
func (r *UserInfoRepo) UserCoins(userID int) (int, error) {
	var coins int
	err := r.db.QueryRow("SELECT coins FROM users WHERE id = $1", userID).Scan(&coins)
	return coins, err
}

// UserInventory возвращает инвентарь пользователя
func (r *UserInfoRepo) UserInventory(userID int) ([]models.Inventory, error) {
	rows, err := r.db.Query("SELECT type, quantity FROM inventory WHERE user_id = $1", userID)
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

	/*
		receivedRows, err := r.db.Query(`
			SELECT from_user, SUM(amount) AS total_amount
			FROM transactions
			WHERE to_user = $1
			GROUP BY from_user`, userID)
	*/

	receivedRows, err := r.db.Query(`
		SELECT users.username, SUM(transactions.amount) AS total_amount
		FROM transactions
		JOIN users ON transactions.from_user =  users.id
		WHERE transactions.to_user = $1
		GROUP BY users.username`, userID)

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

	/*
		sentRows, err := r.db.Query(`
			SELECT to_user, SUM(amount) AS total_amount
			FROM transactions
			WHERE from_user = $1
			GROUP BY to_user`, userID)
	*/

	sentRows, err := r.db.Query(`
		SELECT users.username, SUM(transactions.amount) AS total_amount
		FROM transactions
		JOIN users ON transactions.to_user = users.id
		WHERE transactions.from_user = $1
		GROUP BY users.username`, userID)

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
