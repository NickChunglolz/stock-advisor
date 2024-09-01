package controller

import (
	"github.com/NickChunglolz/stock-advisor/ochestractor-dms/usecase"
	"github.com/gofiber/fiber/v2"
)

type Controllers struct {
	Stock *StockController
}

func InitControllers(usecases *usecase.Usecases) *Controllers {
	return &Controllers{
		Stock: NewStockController(usecases.StockQuery, usecases.StockCommand),
	}
}

func SetupRoutes(app *fiber.App, controllers *Controllers) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/Stocks", fiber.StatusFound)
	})

	controllers.Stock.setupStockInfoRoutes(app)
}
