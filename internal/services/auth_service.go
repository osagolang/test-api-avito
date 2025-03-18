package services

import (
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

// Auth - Аутентификация пользователя
func (s *AuthService) Auth(user *models.User, password string) (string, error) {

	// Проверяем пароль
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	// Генерируем токен
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
