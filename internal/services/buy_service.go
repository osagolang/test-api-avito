package services

import (
	"test-api-avito/internal/repositories"
)

type BuyService struct {
	buyRepo *repositories.BuyRepo
}

func NewBuyService(buyRepo *repositories.BuyRepo) *BuyService {
	return &BuyService{buyRepo: buyRepo}
}
