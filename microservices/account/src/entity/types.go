package entity

import "time"

type Credentials struct {
	Password string `json:"password"`
	Cpf      string `json:"cpf"`
}

type BankAccountResponse struct {
	Msg         string      `json:"message,omitempty"`
	Token       string      `json:"session_token,omitempty"`
	BankAccount BankAccount `json:"bank_account,omitempty"`
}

type Payload struct {
	ID     string `json:"bank_account_id,omitempty" reqHeader:"bank_account_id,omitempty"`
	Status string `json:"status,omitempty"`
	Token  string `json:"session_token,omitempty" reqHeader:"bank_account_id,omitempty"`
	Msg    string `json:"message,omitempty"`
	Cpf    string `json:"cpf,omitempty"`
}

type BankAccount struct {
	ID          string  `json:"bank_account_id"`
	Cpf         string  `json:"cpf"`
	Balance     float64 `json:"balance"`
	BankAccount string  `json:"bank_account"`
	BankAgency  string  `json:"bank_agency"`
	BankCode    string  `json:"bank_code"`
	Status      string  `json:"status"`
	DailyLimit  float64 `json:"daily_limit"`
}

var BankAccounts = map[string]BankAccount{}

var Clock = time.Now()

const URL = "http://localhost:28080/"
