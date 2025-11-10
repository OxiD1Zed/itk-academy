package handler

import (
	"encoding/json"
	"itk-academy/internal/model"
	"net/http"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

func NewWalletHandler(walletService WalletService) *WalletHandler {
	return &WalletHandler{
		walletService: walletService,
	}
}

type WalletHandler struct {
	walletService WalletService
}

type WalletService interface {
	GetBalance(uuid uuid.UUID) (decimal.Decimal, error)
	ChangeBalance(uuid uuid.UUID, operationType model.OperationType, amount decimal.Decimal) error
}

func (h *WalletHandler) HandleChangeBalance(w http.ResponseWriter, r *http.Request) {
	var req ChangeBalanceRequast
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	opType := model.OperationType(strings.ToUpper(req.OperationType))
	if err := h.walletService.ChangeBalance(req.WalletID, opType, req.Amount); err != nil {
		handleErrors(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleErrors(err error, w http.ResponseWriter) {
	switch err {
	case model.ErrorUnknowOperationType:
		http.Error(w, err.Error(), http.StatusBadRequest)
	case model.ErrorClosedConnection:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	case model.ErrorAcquireTimeout:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	case model.ErrorNotFound:
		http.Error(w, err.Error(), http.StatusNotFound)
	case model.ErrorInsufficientFunds:
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		http.Error(w, "unknown error", http.StatusInternalServerError)
	}
}

type ChangeBalanceRequast struct {
	WalletID      uuid.UUID       `json:"walletId"`
	OperationType string          `json:"operationType"`
	Amount        decimal.Decimal `json:"amount"`
}

func (h *WalletHandler) HandleGetBalance(w http.ResponseWriter, r *http.Request) {
	walletID, err := uuid.FromString(r.PathValue("wallet_id"))
	if err != nil {
		http.Error(w, "invalid id format", http.StatusBadRequest)
		return
	}

	balance, err := h.walletService.GetBalance(walletID)
	if err != nil {
		handleErrors(err, w)
		return
	}

	resp := BalanceResponse{
		WalletId: walletID,
		Balance:  balance,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

type BalanceResponse struct {
	WalletId uuid.UUID       `json:"walletId"`
	Balance  decimal.Decimal `json:"balance"`
}
