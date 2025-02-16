package repositories

import (
	"database/sql"
	"test-api-avito/internal/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

// CreateUser добавляет нового пользователя в БД
func (r *UserRepo) CreateUser(username, password string) (*models.User, error) {

	var user models.User

	// Создание нового пользователя
	err := r.db.QueryRow(
		"INSERT INTO users (username, password, coins) VALUES ($1, $2, $3) RETURNING id, username, coins",
		username, password, 1000,
	).Scan(&user.ID, &user.Username, &user.Coins)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindUser проверяет наличие пользователя в БД
func (r *UserRepo) FindUser(username string) (*models.User, error) {

	var user models.User

	err := r.db.QueryRow(
		"SELECT id, username, password FROM users WHERE username = $1",
		username,
	).Scan(&user.ID, &user.Username, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
