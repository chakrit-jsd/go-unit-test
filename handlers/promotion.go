package handlers

import (
	"go-unit-test/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PromotionHandler interface {
	CalculateDiscount(c *fiber.Ctx) error
}

type promotionHandler struct {
	promoSrv services.PromotionService
}

func NewPromotionHandler(promoSrv services.PromotionService) PromotionHandler {
	return promotionHandler{promoSrv}
}

func (h promotionHandler) CalculateDiscount(c *fiber.Ctx) error {

	amountString := c.Query("amount")
	amount, err := strconv.Atoi(amountString)
	if err != nil {
		return fiber.ErrBadRequest
	}

	amount, err = h.promoSrv.CalcutateDiscount(amount)

	if err != nil {
		return fiber.ErrConflict
	}

	return c.JSON(amount)
}
