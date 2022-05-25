package entity

type Payload struct {
	Type          string  `json:"type"`
	Amount        float64 `json:"amount"`
	BankAccountID string  `json:"bank_account_id,omitempty" reqHeader:"bank_account_id,omitempty"`
	Token         string  `json:"session_token,omitempty" reqHeader:"session_token,omitempty"`
	Filter        string  `json:"filter,omitempty" reqHeader:"filter,omitempty"`
	Date          string  `json:"date,omitempty" reqHeader:"date,omitempty"`
	Test          string  `json:"test,omitempty" reqHeader:"test,omitempty"`
}

type BankAccountResponse struct {
	Msg         string      `json:"message,omitempty"`
	Token       string      `json:"session_token,omitempty"`
	BankAccount BankAccount `json:"bank_account,omitempty"`
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

type TransactionResponse struct {
	Msg                string        `json:"message,omitempty"`
	Token              string        `json:"session_token,omitempty"`
	Transaction        []Transaction `json:"transaction,omitempty"`
	TransactionHistory interface{}   `json:"transaction_history,omitempty"`
}

type Transaction struct {
	TransactionID string  `json:"transaction_id"`
	BankAccountID string  `json:"bank_account_id"`
	Type          string  `json:"type"`
	Amount        float64 `json:"amount"`
	CreatedAt     string  `json:"created_at"`
}

var TransactionHistory = map[string][]Transaction{}

var DailyLimit = map[string]float64{}

type Filter struct {
	Date string `json:"date"`
	ID   string `json:"bank_account_id"`
}

var FilterTransactionHistory = map[Filter][]Transaction{}

const SIGNIN_URL = "http://localhost:28080/"
const ACCOUNT_URL = "http://localhost:28081/"
