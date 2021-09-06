package usecase

import (
	"go-mock-example/domain"
	"strconv"
)

type ShopUsecase struct {
	RedisRepository domain.RedisRepository
}

func NewShopUsecase(redisRepository domain.RedisRepository) domain.IShopUsecase {
	return &ShopUsecase{
		RedisRepository: redisRepository,
	}
}

func (s ShopUsecase) GetProductPrice(handle string) (int, error) {
	priceStr, err := s.RedisRepository.Get(handle)
	if err != nil {
		return 0, err
	}

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		return 0, err
	}

	return price, nil
}
