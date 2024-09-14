package main

import (
	"makly/hangman/internal/application"
	"makly/hangman/internal/infrastructure"
)

func main() {
	wordsCollection, category, difficulty, err := infrastructure.ConsoleGameInit()
	if err != nil {
		panic(err) // TODO: handle error with logging
	}
	inputer := infrastructure.NewConsoleInput()
	outputer := infrastructure.NewConsoleOutput()
	application.RunGameSession(wordsCollection, category, difficulty, inputer, outputer)
}
