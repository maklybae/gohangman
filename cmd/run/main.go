package main

import (
	"makly/hangman/internal/application"
	"makly/hangman/internal/infrastructure"

	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("json")   // Look for specific type

	err := viper.ReadInConfig()
	if err != nil {
		panic(err) // TODO: handle error with logging
	}

	category, difficulty, err := infrastructure.ConsoleGameInit(viper.GetString("defaultSamplePath"))
	if err != nil {
		panic(err) // TODO: handle error with logging
	}

	inputer := infrastructure.NewConsoleInput()
	outputer := infrastructure.NewConsoleOutput()

	err = application.RunGameSession(category, difficulty, inputer, outputer)
	if err != nil {
		panic(err) // TODO: handle error with logging
	}
}
