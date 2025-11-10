package app

import (
	"fmt"
	"itk-academy/config"
	"itk-academy/internal/db"
	"itk-academy/internal/handler"
	"itk-academy/internal/router"
	"itk-academy/internal/service"
	"log/slog"
	"net/http"
	"time"

	"github.com/jackc/pgx"
)

func NewApp(log *slog.Logger, config *config.Config) *App {
	return &App{
		log:    log,
		config: config,
	}
}

type App struct {
	log    *slog.Logger
	config *config.Config
}

func (app *App) Run() error {
	const op = "app.App.Run"

	log := app.log.With(
		slog.String("op", op),
	)

	pool, err := createConnectionPool(
		app.config.Postgres.Host,
		app.config.Postgres.Username,
		app.config.Postgres.Password,
		app.config.Postgres.Database,
		app.config.Postgres.Port,
		app.config.Postgres.MaxConnections,
		app.config.Postgres.AcquireTimeout,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	walletProvider := db.NewWalletProvider(*pool)
	walletService := service.NewWalletService(app.log, walletProvider)
	walletHandler := handler.NewWalletHandler(walletService)
	router := router.SetupRoutes(walletHandler)

	log.Info("starting the server")
	defer log.Info("stopping the server")

	add := fmt.Sprintf("%s:%v", app.config.AppConfig.Host, app.config.AppConfig.Port)
	log.Info("server address: %s", add)
	if err := http.ListenAndServe(
		add,
		router,
	); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func createConnectionPool(host, username, password, database string, port uint16, maxConn int, acquireTimeout time.Duration) (*pgx.ConnPool, error) {
	return pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     host,
			Port:     port,
			User:     username,
			Password: password,
			Database: database,
		},
		MaxConnections: maxConn,
		AcquireTimeout: acquireTimeout,
	})
}
