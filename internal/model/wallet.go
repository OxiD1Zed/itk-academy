package model

import (
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

func NewWallet(id uuid.UUID, balance decimal.Decimal) *Wallet {
	return &Wallet{
		id:      id,
		balance: balance,
	}
}

type Wallet struct {
	id      uuid.UUID
	balance decimal.Decimal
}

func (w *Wallet) GetID() uuid.UUID {
	return w.id
}

func (w *Wallet) GetBalance() decimal.Decimal {
	return w.balance
}
