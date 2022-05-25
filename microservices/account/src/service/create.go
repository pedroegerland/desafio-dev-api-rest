package service

import (
	"banktest_account/src/entity"
	"banktest_account/src/helpers"
	"fmt"
	"github.com/gofiber/fiber/v2"
	validator "github.com/paemuri/brdoc"
)

func Create(c *fiber.Ctx) error {
	var credentials entity.Credentials

	if err := c.BodyParser(&credentials); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.Payload{Msg: fmt.Sprintf("Oops! Something Bad Happened! Error: %v", err.Error())})
	}

	if credentials.Cpf == "" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(entity.Payload{Msg: "Please insert your cpf!"})
	}

	if credentials.Password == "" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(entity.Payload{Msg: "Please insert your password!"})
	}

	if !validator.IsCPF(credentials.Cpf) {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(entity.Payload{Msg: "Please insert a valid cpf!"})
	}

	var response entity.Payload
	if err := helpers.ProcessRequest("create", credentials, &response); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.Payload{Msg: fmt.Sprintf("Oops! Something Bad Happened! Error: %v", err.Error())})
	}
	if response.Msg != "User created!" {
		return c.Status(fiber.StatusBadRequest).JSON(entity.Payload{Msg: "User already exists!"})
	}

	token := entity.Payload{}
	if err := helpers.ProcessRequest("signin", credentials, &token); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.Payload{Msg: fmt.Sprintf("Oops! Something Bad Happened! Error: %v", err.Error())})
	}

	bankAccount := helpers.GenerateBankAccount(credentials.Cpf)

	entity.BankAccounts[credentials.Cpf] = bankAccount

	return c.Status(fiber.StatusOK).JSON(entity.BankAccountResponse{BankAccount: bankAccount, Token: token.Token})
}
