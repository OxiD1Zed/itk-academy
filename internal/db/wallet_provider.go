package db

import (
	"itk-academy/internal/model"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx"
	"github.com/shopspring/decimal"
)

func NewWalletProvider(pool pgx.ConnPool) *WalletProvider {
	return &WalletProvider{
		pool: pool,
	}
}

type WalletProvider struct {
	pool pgx.ConnPool
}

func (w *WalletProvider) ChangeBalance(uuid uuid.UUID, amount decimal.Decimal) error {
	sql := "call update_wallet_balance($1, $2)"
	_, err := w.pool.Exec(sql, uuid, amount)
	if err != nil {
		return handleErrors(err)
	}
	return nil
}

func (w *WalletProvider) GetBalance(uuid uuid.UUID) (decimal.Decimal, error) {
	sqlReq := "select get_balance($1)"
	var balance decimal.Decimal
	err := w.pool.QueryRow(sqlReq, uuid).Scan(&balance)
	if err != nil {
		return decimal.Zero, handleErrors(err)
	}

	return balance, nil
}

func handleErrors(err error) error {
	switch err {
	case pgx.ErrAcquireTimeout:
		return model.ErrorAcquireTimeout
	case pgx.ErrClosedPool, pgx.ErrConnBusy, pgx.ErrDeadConn:
		return model.ErrorClosedConnection
	case pgx.ErrNoRows:
		return model.ErrorNotFound
	default:
		return handleDatabaseError(err)
	}
}

func handleDatabaseError(err error) error {
	if pgErr, ok := err.(pgx.PgError); ok {
		switch pgErr.Code {
		case "P0002":
			return model.ErrorNotFound
		case "22P02":
			return model.ErrorInsufficientFunds
		default:
			return err
		}
	}
	return err
}
