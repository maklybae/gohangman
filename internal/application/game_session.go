package application

import (
	"makly/hangman/internal/domain"
)

func RunGameSession(wordsCollection *domain.WordsCollection, category *domain.Category, difficulty domain.Difficulty, inputer domain.GameInputer, outputer domain.GameOutputer) {
	// TODO: check if wordsCollection is nil
	if category == nil {
		category = ChoiceCategory(wordsCollection)
	}
	word := ChoiceWord(category, difficulty)
	game := domain.NewGame(word)
	for !game.IsFinished() {
		outputer.ShowGame(game)
		letter := inputer.GetLetter()
		game.Guess(letter)
	}
	outputer.ShowGameResult(game)
}
