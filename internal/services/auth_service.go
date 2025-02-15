package services

import (
	"golang.org/x/crypto/bcrypt"
	"test-api-avito/internal/models"
	"test-api-avito/internal/repositories"
)

type AuthService struct {
	userRepo   *repositories.UserRepo
	walletRepo *repositories.WalletRepo
}

func NewAuthService(userRepo *repositories.UserRepo, walletRepo *repositories.WalletRepo) *AuthService {
	return &AuthService{userRepo: userRepo, walletRepo: walletRepo}
}

// Register регистрирует нового пользователя и создает кошелёк
func (s *AuthService) Register(username, password string) (*models.User, error) {

	// Хэшируем пароль
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Создаем пользователя через userRepo
	user, err := s.userRepo.CreateUser(username, string(hashPassword))
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Аутентификация пользователя
func (s *AuthService) Login(username, password string) (*models.User, error) {

	// Поиск существующего пользователя
	user, err := s.userRepo.FindUser(username)
	if err != nil {
		return nil, err
	}

	// Проверка пароля
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}
