package services

import (
	"test-api-avito/internal/models"
	"test-api-avito/internal/repositories"
)

type InfoService struct {
	userInfoRepo *repositories.UserInfoRepo
}

func NewInfoService(userInfoRepo *repositories.UserInfoRepo) *InfoService {
	return &InfoService{userInfoRepo: userInfoRepo}
}

// UserInfo возвращает информацию пользователя
func (s *InfoService) UserInfo(userID int) (models.InfoResponse, error) {
	var response models.InfoResponse

	// Получаем количество монет
	coins, err := s.userInfoRepo.UserCoins(userID)
	if err != nil {
		return response, err
	}
	response.Coins = coins

	// Получаем инвентарь
	inventory, err := s.userInfoRepo.UserInventory(userID)
	if err != nil {
		return response, err
	}
	response.Inventory = inventory

	// Получаем историю транзакций
	history, err := s.userInfoRepo.UserTransactionHistory(userID)
	if err != nil {
		return response, err
	}
	response.CoinHistory = history

	return response, nil
}
