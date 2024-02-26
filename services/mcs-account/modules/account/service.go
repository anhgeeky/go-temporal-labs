package account

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	Repo Repository
}

// Sample for test only
func (r Service) GetAccounts() (*[]Account, error) {
	accId1, _ := uuid.Parse("54892431-0a67-4b66-91c7-255d2321b224")
	accId2, _ := uuid.Parse("4ea702d2-b129-483f-b554-e7808fa8b9d7")

	items := []Account{
		{
			Id:        accId1,
			Name:      "Customer 1",
			Cpf:       87832842067,
			Balance:   5123560000,
			CreatedAt: time.Now(),
		},
		{
			Id:        accId2,
			Name:      "Customer 2",
			Cpf:       87832842067,
			Balance:   5123560000,
			CreatedAt: time.Now(),
		},
	}

	return &items, nil
}

// Sample for test only
func (r Service) GetBalance() (*BalanceRes, error) {
	accId1, _ := uuid.Parse("54892431-0a67-4b66-91c7-255d2321b224")

	items := []Account{
		{
			Id:        accId1,
			Name:      "Customer 1",
			Cpf:       87832842067,
			Balance:   5123560000,
			CreatedAt: time.Now(),
		},
	}

	var balance float64
	if len(items) > 0 {
		balance = float64(items[0].Balance)
	}

	res := BalanceRes{}
	res.Balance = balance

	return &res, nil
}
