package services

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"test-api-avito/internal/models"
	"test-api-avito/internal/repositories"
	"test-api-avito/internal/utils"
)

type AuthService struct {
	userRepo *repositories.UserRepo
}

func NewAuthService(userRepo *repositories.UserRepo) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// Login - Авторизация/регистрация пользователя
func (s *AuthService) Login(username, password string) (*models.User, string, error) {

	var user *models.User
	var err error

	// Поиск существующего пользователя
	user, err = s.userRepo.FindUser(username)
	if err != nil {

		// Если не найден, регистрируем
		if err == sql.ErrNoRows {

			// Хэшируем пароль
			hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				return nil, "", err
			}

			// Создаем пользователя через userRepo (там же создаётся кошелёк при создании пользователя)
			user, err = s.userRepo.CreateUser(username, string(hashPassword))
			if err != nil {
				return nil, "", err
			}
		} else {
			return nil, "", err
		}
	} else {

		// Если найден, проверяем пароль
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return nil, "", err
		}
	}

	// Генерируем токен
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
