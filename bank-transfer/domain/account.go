package domain

import (
	"time"

	"github.com/pborman/uuid"
)

// Account content struct deifinition
type Account struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"string"`
	Cpf       int64     `json:"cpf"`
	Secret    string    `json:"-"`
	Balance   float64   `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type AccountList struct {
	Accounts []Account `json:"accounts"`
}

var Accounts = []Account{
	{
		Id:        uuid.Parse("54892431-0a67-4b66-91c7-255d2321b224"),
		Name:      "Customer 1",
		Cpf:       87832842067,
		Balance:   5123560000,
		CreatedAt: time.Now(),
	},
	{
		Id:        uuid.Parse("4ea702d2-b129-483f-b554-e7808fa8b9d7"),
		Name:      "Customer 2",
		Cpf:       87832842067,
		Balance:   5123560000,
		CreatedAt: time.Now(),
	},
}
