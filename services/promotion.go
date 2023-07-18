package services

import (
	"go-unit-test/repositories"
	"math"
)

type PromotionService interface {
	CalcutateDiscount(amount int) (int, error)
}

type promotionService struct {
	promoRepo repositories.PromotionRepository
}

func NewPromotionService(promoRepo repositories.PromotionRepository) PromotionService {
	return promotionService{promoRepo}
}

func (s promotionService) CalcutateDiscount(amount int) (int, error) {

	if amount <= 0 {
		return 0, ErrZeroAmount
	}

	promotion, err := s.promoRepo.GetPromotion()
	if err != nil {
		return 0, ErrRepository
	}

	if amount >= promotion.PurchaseMin {

		return int(math.Round((float64(amount) - (float64(amount) * float64(promotion.DiscountPercent) / 100.0)))), nil
	}

	return amount, nil
}
