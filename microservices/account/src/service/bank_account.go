package service

import (
	"banktest_account/src/entity"
	"banktest_account/src/helpers"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func BankAccount(c *fiber.Ctx) error {
	var payload entity.Payload

	headers := c.GetReqHeaders()

	payload.ID = headers["Bank_account_id"]
	payload.Token = headers["Session_token"]

	token := entity.Payload{Token: payload.Token}

	var response entity.Payload
	if err := helpers.ProcessRequest("validate", token, &response); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.Payload{Msg: fmt.Sprintf("Oops! Something Bad Happened! Error: %v", err.Error())})
	}

	bankAccount, exists := entity.BankAccounts[response.Cpf]
	if !exists {
		return c.Status(fiber.StatusBadRequest).JSON(entity.Payload{Msg: "customer does not have an account", Token: response.Token, Cpf: response.Cpf})
	}

	bankAccount.NeedUpdateLimit(entity.Clock, time.Now())

	return c.Status(fiber.StatusOK).JSON(entity.BankAccountResponse{BankAccount: bankAccount, Token: response.Token})
}
