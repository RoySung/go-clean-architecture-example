package usecase

import (
	"go-mock-example/domain"
	"go-mock-example/domain/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestShopUsecase_GetProductPrice(t *testing.T) {
	c := gomock.NewController(t)
	redisRepoMock := mocks.NewMockRedisRepository(c)
	productHandle := "handle"
	expectPriceStr := "10"
	expectPrice := 10
	redisRepoMock.EXPECT().Get(productHandle).Return(expectPriceStr, nil)

	type fields struct {
		RedisRepository domain.RedisRepository
	}
	type args struct {
		handle string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test ShopUsecase.Get",
			fields{redisRepoMock},
			args{productHandle},
			expectPrice,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ShopUsecase{
				RedisRepository: tt.fields.RedisRepository,
			}
			got, err := s.GetProductPrice(tt.args.handle)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShopUsecase.GetProductPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ShopUsecase.GetProductPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
