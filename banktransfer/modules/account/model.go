package account

import (
	"time"

	"github.com/google/uuid"
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
