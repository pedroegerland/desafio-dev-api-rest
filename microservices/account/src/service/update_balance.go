package service

import (
	"banktest_account/src/entity"
	"banktest_account/src/helpers"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func UpdateBalance(c *fiber.Ctx) error {
	var incomingAccount entity.BankAccountResponse
	if err := c.BodyParser(&incomingAccount); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(entity.Payload{Msg: fmt.Sprintf("Oops! Something Bad Happened! Error: %v", err.Error())})
	}

	token := entity.Payload{Token: incomingAccount.Token}

	var response entity.Payload
	if err := helpers.ProcessRequest("validate", token, &response); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.Payload{Msg: "Oops! Something Bad Happened!"})
	}

	bankAccount, _ := entity.BankAccounts[incomingAccount.BankAccount.Cpf]

	bankAccount = incomingAccount.BankAccount

	entity.BankAccounts[incomingAccount.BankAccount.Cpf] = bankAccount
	return c.Status(fiber.StatusOK).JSON(entity.BankAccountResponse{BankAccount: bankAccount, Token: response.Token})
}
