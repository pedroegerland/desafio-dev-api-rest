package entity

import "time"

func (b BankAccount) IsClosed() bool {
	return b.Status == "closed"
}

func (b BankAccount) HaveBalance() bool {
	return b.Balance < 0 || b.Balance > 0
}

func (b *BankAccount) NeedUpdateLimit(old, new time.Time) {
	if new.Day()-old.Day() >= 1 {
		b.DailyLimit = 2000
		Clock = time.Now()
	}
}
