package services_test

import (
	"errors"
	"go-unit-test/repositories"
	"go-unit-test/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {

	type testCase struct {
		name            string
		purchaseMin     int
		discountPercent int
		amount          int
		expected        int
	}

	cases := []testCase{
		{name: "applied 100", purchaseMin: 100, discountPercent: 20, amount: 100, expected: 80},
		{name: "applied 199", purchaseMin: 100, discountPercent: 30, amount: 199, expected: 139},
		{name: "applied 999", purchaseMin: 100, discountPercent: 99, amount: 999, expected: 10},
		{name: "not applied 99", purchaseMin: 100, discountPercent: 20, amount: 99, expected: 99},
	}

	for _, c := range cases {

		t.Run(c.name, func(t *testing.T) {
			// Arrage
			promoRepo := repositories.NewPromotionRepositoryMock()
			promoRepo.On("GetPromotion").Return(repositories.Promotion{
				ID:              1,
				PurchaseMin:     c.purchaseMin,
				DiscountPercent: c.discountPercent,
			}, nil)

			promoService := services.NewPromotionService(promoRepo)

			// Act
			discount, _ := promoService.CalcutateDiscount(c.amount)
			expected := c.expected
			// Assert
			assert.Equal(t, expected, discount)
		})
	}
	t.Run("amount zero", func(t *testing.T) {
		// Arrage
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{
			ID:              1,
			PurchaseMin:     100,
			DiscountPercent: 20,
		}, nil)

		promoService := services.NewPromotionService(promoRepo)

		// Act
		_, err := promoService.CalcutateDiscount(0)

		// Assert
		assert.ErrorIs(t, err, services.ErrZeroAmount)
		promoRepo.AssertNotCalled(t, "GetPromotion")

	})

	t.Run("repository error", func(t *testing.T) {
		// Arrage
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{}, errors.New("Error"))

		promoService := services.NewPromotionService(promoRepo)

		// Act
		_, err := promoService.CalcutateDiscount(100)

		// Assert
		assert.ErrorIs(t, err, services.ErrRepository)
	})
}
