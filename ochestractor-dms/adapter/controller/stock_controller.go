package controller

import (
	"github.com/NickChunglolz/stock-advisor/ochestractor-dms/usecase"
	"github.com/gofiber/fiber/v2"
)

const (
	domainPath = "/Stocks"
)

func setupStockInfoRoutes(app *fiber.App) {
	app.Get(domainPath, func(c *fiber.Ctx) error {
		data, err := usecase.GetStocks()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(data)
	})
}
