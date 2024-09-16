package main

import (
	"errors"
	"log/slog"
	"makly/hangman/internal/application"
	"makly/hangman/internal/infrastructure"
	"makly/hangman/pkg/climenu"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func main() {
	// Read configuration
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// Initialize logger
	absLogFilePath, err := filepath.Abs(viper.GetString("logPath"))
	if err != nil {
		panic(err)
	}

	logFile, err := os.OpenFile(absLogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	logger := slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{AddSource: true}))
	slog.SetDefault(logger)

	// Initialize game
	category, difficulty, err := infrastructure.Init(viper.GetString("defaultSamplePath"), viper.GetString("jsonSchemaPath"))
	if err != nil {
		var exitErr *climenu.ExitError
		if errors.As(err, &exitErr) {
			slog.Info("Game session ended", slog.String("reason", err.Error()))
			return
		}

		slog.Error("Initialization error", slog.Any("error", err))
		panic(err)
	}

	inputer := infrastructure.NewConsoleInput()
	outputer := infrastructure.NewConsoleOutput()

	// Run game session
	if err := application.RunGameSession(category, difficulty, inputer, outputer); err != nil {
		slog.Error("Game session error", slog.Any("error", err))
		panic(err)
	}

	slog.Info("Game session ended", slog.String("reason", "success"))
}
