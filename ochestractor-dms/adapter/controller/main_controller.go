package controller

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/Stocks", fiber.StatusFound)
	})

	setupStockInfoRoutes(app)
}
