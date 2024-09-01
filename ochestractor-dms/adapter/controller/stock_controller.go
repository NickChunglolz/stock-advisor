package controller

import (
	"github.com/NickChunglolz/stock-advisor/ochestractor-dms/usecase"
	"github.com/gofiber/fiber/v2"
)

const (
	domainPath = "/Stocks"
)

type StockController struct {
	query   *usecase.StockQuery
	command *usecase.StockCommand
}

func NewStockController(query *usecase.StockQuery, command *usecase.StockCommand) *StockController {
	return &StockController{
		query:   query,
		command: command,
	}
}

func (controller *StockController) setupStockInfoRoutes(app *fiber.App) {
	app.Get(domainPath, func(c *fiber.Ctx) error {
		data, err := controller.query.GetStocks()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(data)
	})
}
