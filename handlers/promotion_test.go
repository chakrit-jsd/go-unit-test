package handlers_test

import (
	"errors"
	"fmt"
	"go-unit-test/handlers"
	"go-unit-test/services"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrage
		amount := 100
		expected := 80

		promoServiceMock := services.NewPromotionServiceMock()
		promoServiceMock.On("CalcutateDiscount", amount).Return(expected, nil)

		promoHandler := handlers.NewPromotionHandler(promoServiceMock)
		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		// Act
		res, _ := app.Test(req)

		// Assert
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(expected), string(body))
		}
	})

	t.Run("fail amount not number", func(t *testing.T) {
		// Arrage
		amount := "xx"
		expected := fiber.ErrBadRequest

		promoServiceMock := services.NewPromotionServiceMock()
		promoServiceMock.On("CalcutateDiscount", amount).Return(nil, errors.New("some text"))

		promoHandler := handlers.NewPromotionHandler(promoServiceMock)
		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		// Act
		res, _ := app.Test(req)

		// Assert
		assert.Equal(t, expected.Code, res.StatusCode)
	})
	t.Run("fail calculate service", func(t *testing.T) {
		// Arrage
		amount := 0
		expected := fiber.ErrConflict

		promoServiceMock := services.NewPromotionServiceMock()
		promoServiceMock.On("CalcutateDiscount", amount).Return(0, errors.New("some text"))

		promoHandler := handlers.NewPromotionHandler(promoServiceMock)
		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		// Act
		res, _ := app.Test(req)

		// Assert
		assert.Equal(t, expected.Code, res.StatusCode)
	})

}
