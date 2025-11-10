package router

import (
	"itk-academy/internal/handler"
	"net/http"
)

func SetupRoutes(walletHandler *handler.WalletHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/wallet", walletHandler.HandleChangeBalance)
	mux.HandleFunc("GET /api/v1/wallets/{wallet_id}", walletHandler.HandleGetBalance)

	return mux
}
