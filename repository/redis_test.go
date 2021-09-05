package repository

import (
	"errors"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

func TestRedisStore_Get(t *testing.T) {
	key := "test-key"
	expectValue := "test-value"
	mockClient, mock := redismock.NewClientMock()
	mock.ExpectGet(key).SetVal(expectValue)

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
			fields{mockClient},
			args{key},
			expectValue,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := RedisRepository{
				Client: tt.fields.Client,
			}
			got, err := s.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("RedisStore.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Error(err)
			}

			assert.Equal(t, got, expectValue)
		})
	}
}

func TestRedisStore_Set(t *testing.T) {
	key := "test-key123"
	expectValue := "test-value"
	mockClient, mock := redismock.NewClientMock()
	testError := errors.New("FAIL")
	mock.ExpectSet(key, expectValue, 0).SetErr(testError)

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
		wantErr error
	}{
		// TODO: Add test cases.
		{
			"test Set method",
			fields{mockClient},
			args{key, expectValue},
			testError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := RedisRepository{
				Client: tt.fields.Client,
			}
			if err := s.Set(tt.args.key, tt.args.value); err != tt.wantErr {
				t.Errorf("RedisStore.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Error(err)
			}
		})
	}
}
