package service

import (
	"github.com/gofiber/fiber/v2"
)

func Health(c *fiber.Ctx) error {
	return c.SendString("Ok")
}

func Liveness(c *fiber.Ctx) error {
	return c.SendString("Ok")
}
