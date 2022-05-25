package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

func (b BankAccount) ExceedsDailyLimit(amount float64) bool {
	return b.DailyLimit < amount
}

func (b BankAccount) LeftoversDailyLimit(amount float64) float64 {
	b.DailyLimit -= amount
	return b.DailyLimit
}

func (b BankAccount) HaveBalanceToTransact(amount float64) bool {
	return b.Balance >= amount
}

func (b BankAccount) Transact(t Transaction) float64 {
	if t.Type == "withdraw" {
		b.Balance -= t.Amount
	}

	if t.Type == "deposit" {
		b.Balance += t.Amount
	}
	return b.Balance
}

func (b *BankAccount) ExecuteTransaction(t Transaction, typeTransaction string) error {
	if typeTransaction == "withdraw" {
		if b.ExceedsDailyLimit(t.Amount) {
			return errors.New("transaction amount exceeds daily limit")
		}
		if !b.HaveBalanceToTransact(t.Amount) {
			return errors.New("transaction amount exceeds balance")
		}
		b.DailyLimit = b.LeftoversDailyLimit(t.Amount)
	}
	b.Balance = b.Transact(t)
	return nil
}

func (t Transaction) InsertInTransactionHistory(f Filter) {
	TransactionHistory[f.ID] = append(TransactionHistory[f.ID], t)
	FilterTransactionHistory[f] = append(FilterTransactionHistory[f], t)
}

func (t Transaction) GetTransactionHistory(id string, f Filter) []Transaction {
	if f.ID != "" {
		return FilterTransactionHistory[f]
	}
	return TransactionHistory[id]
}

func (p Payload) EventDate(f bool) string {
	t := time.Now()
	if p.Test == "test" {
		t = t.Add(-24 * time.Hour)
	}

	newDate := t.Format("2006-01-02")
	if f {
		newDate = t.Format("2006-01-02T15:04:05.000Z")
	}
	return newDate
}
func (p Payload) GenerateTransaction() Transaction {
	t := p.EventDate(true)
	return Transaction{
		TransactionID: uuid.NewString(),
		BankAccountID: p.BankAccountID,
		Type:          p.Type,
		Amount:        p.Amount,
		CreatedAt:     t,
	}
}
