package service

import (
	"itk-academy/internal/model"
	"log/slog"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

func NewWalletService(log *slog.Logger, walletProvider WalletProvider) *WalletService {
	return &WalletService{
		log:            log,
		walletProvider: walletProvider,
	}
}

type WalletProvider interface {
	GetBalance(uuid uuid.UUID) (decimal.Decimal, error)
	ChangeBalance(uuid uuid.UUID, amount decimal.Decimal) error
}

type WalletService struct {
	log            *slog.Logger
	walletProvider WalletProvider
}

func (ws *WalletService) GetBalance(uuid uuid.UUID) (decimal.Decimal, error) {
	const op = "service.WalletService.GetBalance"

	log := ws.log.With(
		slog.String("op", op),
	)

	log.Info("uuid: %s: the beginning of receiving the balance", uuid)

	balance, err := ws.walletProvider.GetBalance(uuid)
	if err != nil {
		log.Error("couldn't get balance", slog.String("error", err.Error()))
		return decimal.Zero, err
	}

	log.Info("uuid: %s: balance successfully received", uuid)

	return balance, nil
}

func (ws *WalletService) ChangeBalance(uuid uuid.UUID, operationType model.OperationType, amount decimal.Decimal) error {
	const op = "service.WalletService.ChangeBalance"

	log := ws.log.With(
		slog.String("op", op),
	)

	log.Info("uuid: %s: the beginning of the balance change", uuid)

	var err error
	switch operationType {
	case model.OperationDeposit:
		err = ws.walletProvider.ChangeBalance(uuid, amount)
	case model.OperationWithDraw:
		amount = amount.Neg()
		err = ws.walletProvider.ChangeBalance(uuid, amount)
	default:
		err = model.ErrorUnknowOperationType
	}

	if err != nil {
		log.Error("couldn't update balance", slog.String("error", err.Error()))
		return err
	}

	log.Info("uuid: %s: the balance has been successfully changed to %v", uuid, amount)

	return nil
}
