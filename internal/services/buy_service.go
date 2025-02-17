package services

import (
	"errors"
	"test-api-avito/internal/repositories"
)

type BuyService struct {
	userInfoRepo *repositories.UserInfoRepo
	itemRepo     *repositories.ItemRepo
}

func NewBuyService(userInfoRepo *repositories.UserInfoRepo, itemRepo *repositories.ItemRepo) *BuyService {
	return &BuyService{userInfoRepo: userInfoRepo, itemRepo: itemRepo}
}

// Покупка предмета
func (s *BuyService) BuyItem(userID int, itemName string) error {
	// Получаем товар
	item, err := s.itemRepo.GetItem(itemName)
	if err != nil {
		return errors.New("товар не найден")
	}

	// Начинаем транзакцию
	tx, err := s.userInfoRepo.BeginTransaction()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Проверяем баланс пользователя
	var userCoins int
	err = tx.QueryRow("SELECT coins FROM users WHERE id = $1 FOR UPDATE", userID).Scan(&userCoins)
	if err != nil {
		return err
	}
	if userCoins < item.Price {
		return errors.New("недостаточно монет")
	}

	// Вычитаем монеты
	_, err = tx.Exec("UPDATE users SET coins = coins - $1 WHERE id = $2", item.Price, userID)
	if err != nil {
		return err
	}

	// Добавляем товар в инвентарь
	_, err = tx.Exec("INSERT INTO inventory (user_id, item, quantity) VALUES ($1, $2, 1) ON CONFLICT (user_id, item) DO UPDATE SET quantity = inventory.quantity + 1", userID, item.Type)
	if err != nil {
		return err
	}

	// Фиксируем транзакцию
	return tx.Commit()
}
