package repositories

import (
	"database/sql"
)

type WalletRepo struct {
	db *sql.DB
}

// NewWalletRepo создает новый экземпляр WalletRepo.
func NewWalletRepo(db *sql.DB) *WalletRepo {
	return &WalletRepo{db: db}
}

// GetBalance возвращает баланс пользователя
func (r *WalletRepo) GetBalance(userID int) (int, error) {

	var balance int

	err := r.db.QueryRow("SELECT balance FROM wallet WHERE user_id = $1", userID).Scan(&balance)

	return balance, err
}
