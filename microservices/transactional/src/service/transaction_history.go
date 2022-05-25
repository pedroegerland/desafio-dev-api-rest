package service

import (
	"banktest_transactional/src/entity"
	"banktest_transactional/src/helpers"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func TransactionHistory(c *fiber.Ctx) error {
	var payload entity.Payload
	var transactionHistory interface{}

	headers := c.GetReqHeaders()

	payload.BankAccountID = headers["Bank_account_id"]
	payload.Token = headers["Session_token"]
	payload.Date = headers["Date"]

	token := entity.Payload{Token: payload.Token}

	var response entity.Payload
	if err := helpers.ProcessRequest("POST", entity.SIGNIN_URL+"validate", "", "", token, &response); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.TransactionResponse{Msg: fmt.Sprintf("Oops! Something Bad Happened! Error: %v", err.Error()), Token: payload.Token})
	}

	transactionHistory = entity.TransactionHistory[payload.BankAccountID]

	if payload.Date != "" {
		filter := entity.Filter{ID: payload.BankAccountID, Date: payload.Date}
		transactionHistory = entity.FilterTransactionHistory[filter]
	}
	return c.Status(fiber.StatusOK).JSON(entity.TransactionResponse{TransactionHistory: transactionHistory, Token: response.Token})
}
