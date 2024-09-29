package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"makly/hangman/internal/application"
	"makly/hangman/internal/infrastructure"
	"makly/hangman/pkg/climenu"

	"github.com/spf13/viper"
)

func main() {
	// Read configuration
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file", err)
		os.Exit(1)
	}

	// Initialize logger
	absLogFilePath, err := filepath.Abs(viper.GetString("logPath"))
	if err != nil {
		fmt.Println("Error getting absolute path os logger", err)
		os.Exit(1)
	}

	logFile, err := os.OpenFile(absLogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		fmt.Println("Error opening log file", err)
		os.Exit(1)
	}

	logger := slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{AddSource: true}))
	slog.SetDefault(logger)

	// Initialize game
	category, difficulty, maxMistakes, err := infrastructure.Init(viper.GetString("defaultSamplePath"), viper.GetString("jsonSchemaPath"))
	if err != nil {
		var exitErr *climenu.ExitError
		if errors.As(err, &exitErr) {
			slog.Info("Game session ended", slog.String("reason", err.Error()))
			return
		}

		slog.Error("Initialization error", slog.Any("error", err))
		logFile.Close()
		os.Exit(1)
	}

	inputer := infrastructure.NewConsoleInput()
	outputer := infrastructure.NewConsoleOutput()

	// Run game session
	randDefault := &application.RandomDefault{}
	if err := application.RunGameSession(category, difficulty, maxMistakes, inputer, outputer, randDefault); err != nil {
		slog.Error("Game session error", slog.Any("error", err))
		logFile.Close()
		os.Exit(1)
	}

	slog.Info("Game session ended", slog.String("reason", "success"))
	logFile.Close()
}
