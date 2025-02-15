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

// CreateUser добавляет нового пользователя в БД и добавляет для него кошелёк
func (r *UserRepo) CreateUser(username, password string) (*models.User, error) {

	var user models.User

	// Начало транзакции (т.к. запись будет сразу в две таблицы, чтобы всё ок было)
	tx, err := r.db.Begin()
	if err != nil {
		return &user, err
	}

	// Создание нового пользователя
	err = tx.QueryRow(
		"INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id, username",
		username, password,
	).Scan(&user.ID, &user.Username)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Создание кошелька с привязкой к ID в users
	_, err = tx.Exec("INSERT INTO wallet (user_id, balance) VALUES ($1, $2)", user.ID, 1000)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Если создание пользователя и кошелька прошло без ошибок, фиксируем транзакцию
	if err := tx.Commit(); err != nil {
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
