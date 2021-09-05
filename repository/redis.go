package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
)

//go:generate mockgen -source=redis.go -destination=mocks/mock_repository.go -package=mocks
type RedisRepository struct {
	Client *redis.Client
}

func (s RedisRepository) Get(key string) (string, error) {
	result, err := s.Client.Get(context.Background(), key).Result()

	return result, err
}

func (s RedisRepository) Set(key string, value string) error {
	err := s.Client.Set(context.Background(), key, value, 0).Err()

	return err
}

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
