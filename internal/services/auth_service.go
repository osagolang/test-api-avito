package services

import (
	"golang.org/x/crypto/bcrypt"
	"test-api-avito/internal/models"
	"test-api-avito/internal/repositories"
)

type AuthService struct {
	userRepo *repositories.UserRepo
}

func NewAuthService(userRepo *repositories.UserRepo) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// Register - Регистрация нового пользователя
func (s *AuthService) Register(username, password string) (*models.User, error) {

	// Хэшируем пароль (уменьшил Cost до 4 для снижения времени ответа)
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	// Создаем пользователя
	user, err := s.userRepo.CreateUser(username, string(hashPassword))
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login - Аутентификация пользователя
func (s *AuthService) Login(username, password string) (*models.User, error) {

	// Поиск существующего пользователя
	user, err := s.userRepo.FindUser(username)
	if err != nil {
		return nil, err
	}

	// Проверяем пароль
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}
