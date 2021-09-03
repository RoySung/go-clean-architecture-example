package main

import (
	"fmt"
	"go-mock-example/repository"
)

func main() {
	store := repository.RedisStore{
		Client: repository.RedisClient,
	}
	key := "test-key"
	store.Set(key, "test-value")
	result, _ := store.Get(key)

	fmt.Printf("%+v\n", result)
}
