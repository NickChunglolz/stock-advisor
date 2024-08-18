package usecase

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var count int

func Increment(c *fiber.Ctx) error {
	count++
	return c.SendString(fmt.Sprintf("%d", count))
}

func Decrement(c *fiber.Ctx) error {
	count--
	return c.SendString(fmt.Sprintf("%d", count))
}
