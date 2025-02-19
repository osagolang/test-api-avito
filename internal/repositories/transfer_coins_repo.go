package repositories

import (
	"database/sql"
	"errors"
)

type TransferCoinsRepo struct {
	db *sql.DB
}

// Перевод монет (с транзакцией)
func (r *TransferCoinsRepo) TransferCoins(fromUserID, toUserID, amount int) error {
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
