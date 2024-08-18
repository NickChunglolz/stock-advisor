package controller

import (
	views "github.com/NickChunglolz/stock-advisor/ochestractor-ams/views"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/overview", fiber.StatusFound)
	})

	setupOverviewRoutes(app)
	setupCounterRoutes(app)

	app.Use(NotFoundMiddleware)
}

func NotFoundMiddleware(c *fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	return Render(c, views.NotFound())
}

func Render(c *fiber.Ctx, content templ.Component) error {
	c.Set("Content-Type", "text/html")
	return views.Index(content).Render(c.Context(), c.Response().BodyWriter())
}
