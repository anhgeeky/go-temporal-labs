package account

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	Repo Repository
}

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
