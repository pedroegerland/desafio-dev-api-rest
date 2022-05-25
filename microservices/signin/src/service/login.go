package service

import (
	"banktest_signin/src/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

func SignIn(c *fiber.Ctx) error {
	var credential entity.Credentials

	if err := c.BodyParser(&credential); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	expectedPassword, exists := entity.Users[credential.Cpf]

	if !exists || expectedPassword != credential.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(entity.Response{Msg: "Unauthorized"})
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(15 * time.Minute).UTC()

	entity.Sessions[sessionToken] = entity.Session{
		Cpf:    credential.Cpf,
		Expiry: expiresAt,
	}

	c.Cookie(&fiber.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})
	return c.Status(fiber.StatusOK).JSON(entity.Response{Cpf: credential.Cpf, SessionToken: sessionToken, ExpiresAt: expiresAt.Format("2006-01-02T15:04:05.000Z")})
}
