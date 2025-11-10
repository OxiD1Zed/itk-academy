package db

import (
	"itk-academy/config"
	"itk-academy/internal/db"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
)

func setup() (*db.WalletProvider, error) {
	if err := godotenv.Load("../config/config_test.env"); err != nil {
		panic("couldn't read the config")
	}
	config := config.NewConfig()
	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     config.Postgres.Host,
			Port:     config.Postgres.Port,
			User:     config.Postgres.Username,
			Password: config.Postgres.Password,
			Database: config.Postgres.Database,
		},
		MaxConnections: config.Postgres.MaxConnections,
		AcquireTimeout: config.Postgres.AcquireTimeout,
	})
	if err != nil {
		return nil, err
	}
	return db.NewWalletProvider(*pool), nil
}

func TestGetBalance(t *testing.T) {
	walletProvider, err := setup()
	if err != nil {
		t.Error(err)
	}

	testWallet := uuid.FromStringOrNil("410a4e9b-45a5-49f4-86d0-dd3af8a1c430")

	_, err = walletProvider.GetBalance(testWallet)
	if err != nil {
		t.Error()
	}
}

func TestChangeBalance(t *testing.T) {
	walletProvider, err := setup()
	if err != nil {
		t.Error(err)
	}

	testWallet := uuid.FromStringOrNil("410a4e9b-45a5-49f4-86d0-dd3af8a1c430")
	testAmount := decimal.NewFromInt(100)

	err = walletProvider.ChangeBalance(testWallet, testAmount)
	if err != nil {
		t.Error(err)
	}
}
