package main

import (
	"itk-academy/config"
	"itk-academy/internal/app"
	"log/slog"
	"os"
)

func main() {
	config := config.NewConfig()
	log := slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug},
		),
	)

	app := app.NewApp(log, config)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
