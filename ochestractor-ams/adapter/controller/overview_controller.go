package controller

import (
	views "github.com/NickChunglolz/stock-advisor/ochestractor-ams/views"
	"github.com/gofiber/fiber/v2"
)

func setupOverviewRoutes(app *fiber.App) {
	app.Get("/overview", func(c *fiber.Ctx) error {
		return Render(c, views.Overview())
	})
}
