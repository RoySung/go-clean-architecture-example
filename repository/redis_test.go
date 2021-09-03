package repository

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
)

var redisMockService *miniredis.Miniredis
var redisClient *redis.Client

func setUp() {
	redisMockService = mockRedis()
	redisClient = redis.NewClient(&redis.Options{
		Addr: redisMockService.Addr(),
	})
}

func teardown() {
	redisMockService.Close()
}

func mockRedis() *miniredis.Miniredis {
	s, err := miniredis.Run()

	if err != nil {
		panic(err)
	}

	return s
}

func TestRedisStore_Get(t *testing.T) {
	setUp()
	defer teardown()
	key := "test-key123"
	expectValue := "test-value"
	redisClient.Set(context.Background(), key, expectValue, 0)

	type fields struct {
		Client *redis.Client
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test get method",
			fields{redisClient},
			args{key},
			expectValue,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := RedisStore{
				Client: tt.fields.Client,
			}
			got, err := s.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("RedisStore.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RedisStore.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisStore_Set(t *testing.T) {
	setUp()
	defer teardown()
	key := "test-key123"
	expectValue := "test-value"

	type fields struct {
		Client *redis.Client
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test Set method",
			fields{redisClient},
			args{key, expectValue},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := RedisStore{
				Client: tt.fields.Client,
			}
			if err := s.Set(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("RedisStore.Set() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				if value, _ := s.Get(tt.args.key); value != expectValue {
					t.Errorf("RedisStore.Set() value = %v, wantValue %v", value, expectValue)
				}
			}
		})
	}
}
