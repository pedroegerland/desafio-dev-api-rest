package service

import (
	"banktest_account/src/entity"
	"banktest_account/src/helpers"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Status(c *fiber.Ctx) error {
	var token entity.Payload
	if err := c.BodyParser(&token); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.Payload{Msg: fmt.Sprintf("Oops! Something Bad Happened! Error: %v", err.Error())})
	}

	if token.Token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(entity.Payload{Msg: "Unauthorized"})
	}

	var response entity.Payload
	if err := helpers.ProcessRequest("validate", token, &response); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.Payload{Msg: fmt.Sprintf("Oops! Something Bad Happened! Error: %v", err.Error())})
	}

	bankAccount, exists := entity.BankAccounts[response.Cpf]
	if !exists {
		return c.Status(fiber.StatusBadRequest).JSON(entity.Payload{Msg: "customer does not have an account", Token: response.Token, Cpf: response.Cpf})
	}

	if bankAccount.IsClosed() {
		return c.Status(fiber.StatusBadRequest).JSON(entity.Payload{Msg: "Account is Closed", Token: response.Token, Cpf: bankAccount.Cpf})
	}

	if token.Status == bankAccount.Status {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.Payload{Msg: fmt.Sprintf("Account already have state %v!", token.Status), Token: response.Token, Cpf: response.Cpf})
	}

	if token.Status == "closed" {
		if bankAccount.HaveBalance() {
			return c.Status(fiber.StatusInternalServerError).JSON(entity.Payload{Msg: fmt.Sprintf("Account still have balance, please withdraw value to continue closing account! Balance: %v", bankAccount.Balance), Token: response.Token, Cpf: bankAccount.Cpf})
		}
	}

	bankAccount.Status = token.Status

	entity.BankAccounts[response.Cpf] = bankAccount

	return c.Status(fiber.StatusOK).JSON(entity.Payload{Msg: fmt.Sprintf("Account %v!", token.Status), Token: response.Token, Cpf: bankAccount.Cpf})
}
