package service

import (
	"banktest_signin/src/entity"
	"github.com/gofiber/fiber/v2"
	"time"
)

func SignOut(c *fiber.Ctx) error {
	sessionToken := c.Cookies("session_token", "")
	if sessionToken == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	delete(entity.Sessions, sessionToken)

	c.Cookie(&fiber.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	})
	return c.Status(fiber.StatusOK).JSON(entity.Response{Msg: "Sign Out with Success"})
}
