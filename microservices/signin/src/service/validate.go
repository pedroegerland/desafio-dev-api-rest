package service

import (
	"banktest_signin/src/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

func Validate(c *fiber.Ctx) error {
	var token entity.Token
	if err := c.BodyParser(&token); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userSession, exists := entity.Sessions[token.SessionToken]
	if !exists {
		return c.Status(fiber.StatusUnauthorized).JSON(entity.Response{Msg: "Unauthorized", SessionToken: token.SessionToken})
	}

	if userSession.IsExpired() {
		newSessionToken := uuid.NewString()
		expiresAt := time.Now().Add(15 * time.Minute).UTC()

		entity.Sessions[newSessionToken] = entity.Session{
			Cpf:    userSession.Cpf,
			Expiry: expiresAt,
		}

		delete(entity.Sessions, token.SessionToken)
		return c.Status(fiber.StatusOK).JSON(entity.Response{Cpf: userSession.Cpf, SessionToken: newSessionToken, ExpiresAt: expiresAt.Format("2006-01-02T15:04:05.000Z")})
	}
	return c.Status(fiber.StatusOK).JSON(entity.Response{Msg: "token is active, there is no need in refresh it", Cpf: userSession.Cpf, SessionToken: token.SessionToken})
}
