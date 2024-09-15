package main

import (
	"errors"
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

	category, difficulty, err := infrastructure.Init(viper.GetString("defaultSamplePath"), viper.GetString("jsonSchemaPath"))
	if err != nil {
		var incorrectJSONError *infrastructure.IncorrectJSONError
		if errors.As(err, &incorrectJSONError) {

		}
		panic(err) // TODO: handle error with logging
	}

	inputer := infrastructure.NewConsoleInput()
	outputer := infrastructure.NewConsoleOutput()

	err = application.RunGameSession(category, difficulty, inputer, outputer)
	if err != nil {
		panic(err) // TODO: handle error with logging
	}
}
