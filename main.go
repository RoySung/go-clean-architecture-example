package main

import (
	"fmt"
	"go-mock-example/repository"
	"go-mock-example/usecase"
)

func main() {
	redisRepository := &repository.RedisRepository{
		Client: repository.RedisClient,
	}
	shopUsecase := usecase.NewShopUsecase(redisRepository)
	productName := "iPhone 13"
	productPrice, err := shopUsecase.GetProductPrice(productName)
	if err != nil {
		fmt.Printf("Error: %+v", err.Error())
		return
	}

	fmt.Printf("%s Price: %d", productName, productPrice)
}
