package main

import (
	"makly/hangman/internal/application"
	"makly/hangman/internal/infrastructure"
)

func main() {
	category, difficulty, err := infrastructure.ConsoleGameInit()
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
