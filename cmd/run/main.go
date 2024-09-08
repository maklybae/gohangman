package main

import (
	"makly/hangman/internal/application"
	"makly/hangman/internal/domain"
	"makly/hangman/internal/infrastructure"
)

func main() {
	consoleInput := &infrastructure.ConsoleInput{}
	consoleOutput := &infrastructure.ConsoleOutput{}
	wordsCollection, err := infrastructure.ReadCollectionFromFile("/Users/makly/Programming/tinkoff-backend-academy/backend-academy_2024_project_1-go-maklybae-4/sample.json")
	if err != nil {
		panic(err)
	}
	application.RunGameSession(wordsCollection, nil, domain.EasyDifficulty, consoleInput, consoleOutput)
}
