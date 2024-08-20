package controller

import (
	"github.com/NickChunglolz/stock-advisor/ochestractor-ams/adapter/client"
	views "github.com/NickChunglolz/stock-advisor/ochestractor-ams/views"
	"github.com/gofiber/fiber/v2"
)

func setupOverviewRoutes(app *fiber.App) {
	app.Get("/overview", func(c *fiber.Ctx) error {
		return Render(c, views.Overview())
	})

	app.Get("/Stocks", func(c *fiber.Ctx) error {
		apiClient := client.NewStockClient()

		data, err := apiClient.GetStocks()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error fetching stock data")
		}

		return c.JSON(data)
	})
}
