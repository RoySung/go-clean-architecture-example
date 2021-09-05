package usecase

import (
	"go-mock-example/domain/mocks"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("ShopGinkgo", func() {
	var (
		mockCtrl      *gomock.Controller
		redisRepoMock *mocks.MockRedisRepository
		shopUsecase   *ShopUsecase
	)
	const productHandle = "product-handle"
	const expectPriceStr = "10"
	const expectPrice = 10

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		redisRepoMock = mocks.NewMockRedisRepository(mockCtrl)
		shopUsecase = NewShopUsecase(redisRepoMock)
	})

	It("ShopUsecase.GetProductPrice", func() {
		redisRepoMock.EXPECT().Get(productHandle).Return(expectPriceStr, nil)
		price, err := shopUsecase.GetProductPrice(productHandle)

		assert.Equal(GinkgoT(), expectPrice, price)
		assert.Equal(GinkgoT(), err, "")
		// Expect(price).To(Equal(expectPrice))
		Expect(err).To(BeNil())
	})
})
