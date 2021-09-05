package domain

//go:generate mockgen -source=redis.go -destination=mocks/mock_redis.go -package=mocks
type RedisRepository interface {
	Get(string) (string, error)
	Set(key string, value string) error
}
