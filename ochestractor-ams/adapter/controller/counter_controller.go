package controller

import (
	"github.com/NickChunglolz/stock-advisor/ochestractor-ams/usecase"
	views "github.com/NickChunglolz/stock-advisor/ochestractor-ams/views"
	"github.com/gofiber/fiber/v2"
)

func setupCounterRoutes(app *fiber.App) {

	app.Get("/counter", func(c *fiber.Ctx) error {
		return Render(c, views.Counter())
	})

	app.Post("/increment", usecase.Increment)
	app.Post("/decrement", usecase.Decrement)
}
