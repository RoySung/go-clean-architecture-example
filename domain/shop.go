package domain

type IShopUsecase interface {
	GetProductPrice(handle string) (int, error)
}
