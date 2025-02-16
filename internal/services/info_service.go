package services

import (
	"test-api-avito/internal/repositories"
)

type InfoService struct {
	userRepo *repositories.UserRepo
}

func NewInfoService(userRepo *repositories.UserRepo) *InfoService {
	return &InfoService{userRepo: userRepo}
}

// UserInfo возвращает информацию пользователя (пока только количество монет)
func (s *InfoService) UserInfo(userID int) (int, error) {

	coins, err := s.userRepo.UserCoins(userID)
	if err != nil {
		return 0, err
	}
	return coins, nil
}
