package service

import (
	"banktest_transactional/src/entity"
	"banktest_transactional/src/helpers"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Transaction(c *fiber.Ctx) error {
	var payload entity.Payload
	var err error

	headers := c.GetReqHeaders()

	if err = c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(entity.TransactionResponse{Msg: fmt.Sprintf("Oops! Something Bad Happened! Error: %v", err.Error()), Token: payload.Token})
	}

	payload.Token = headers["Session_token"]
	payload.BankAccountID = headers["Bank_account_id"]
	payload.Test = headers["Test"]

	token := entity.Payload{Token: payload.Token}

	var response entity.Payload
	if err = helpers.ProcessRequest("POST", entity.SIGNIN_URL+"validate", "", "", token, &response); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.TransactionResponse{Msg: fmt.Sprintf("Oops! Something Bad Happened! Error: %v", err.Error()), Token: response.Token})
	}

	var bankAccountResponse entity.BankAccountResponse
	if err = helpers.ProcessRequest("GET", entity.ACCOUNT_URL+"account", payload.Token, payload.BankAccountID, token, &bankAccountResponse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.TransactionResponse{Msg: fmt.Sprintf("Oops! Something Bad Happened! Error: %v", err.Error()), Token: response.Token})
	}

	if bankAccountResponse.BankAccount.Status != "enabled" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(entity.TransactionResponse{Msg: "Oops! Something Bad Happened!", Token: response.Token, Transaction: []entity.Transaction{{BankAccountID: bankAccountResponse.BankAccount.ID}}})
	}

	transaction := payload.GenerateTransaction()

	bankAccountResponse.BankAccount, err = bankAccountResponse.BankAccount.ExecuteTransaction(transaction, payload.Type)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.TransactionResponse{Msg: fmt.Sprintf("Oops! Something Bad Happened! Error: %v", err.Error()), Token: response.Token, Transaction: []entity.Transaction{{BankAccountID: bankAccountResponse.BankAccount.ID}}})
	}

	var respUpdateBalance entity.BankAccountResponse
	if err = helpers.ProcessRequest("POST", entity.ACCOUNT_URL+"balance", "", "", bankAccountResponse, &respUpdateBalance); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.TransactionResponse{Msg: fmt.Sprintf("Oops! Something Bad Happened! Error: %v", err.Error()), Token: response.Token, Transaction: []entity.Transaction{{BankAccountID: bankAccountResponse.BankAccount.ID}}})
	}

	filter := entity.Filter{ID: payload.BankAccountID, Date: payload.EventDate(false)}
	transaction.InsertInTransactionHistory(filter)

	return c.Status(fiber.StatusOK).JSON(entity.TransactionResponse{Transaction: []entity.Transaction{transaction}, Token: response.Token})
}
