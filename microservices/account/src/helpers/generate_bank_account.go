package helpers

import (
	"banktest_account/src/entity"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
)

func GenerateBankAccount(cpf string) entity.BankAccount {

	return entity.BankAccount{
		ID:          uuid.NewString(),
		Cpf:         cpf,
		Balance:     0,
		BankAgency:  strconv.Itoa(rand.Intn(9999-1000) + 1000),
		BankAccount: strconv.Itoa(rand.Intn(999999999-100000000) + 100000000),
		BankCode:    "123",
		Status:      "enabled",
		DailyLimit:  2000,
	}
}
