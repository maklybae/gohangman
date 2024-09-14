package application

import (
	"fmt"
	"makly/hangman/internal/domain"
)

func RunGameSession(
	category *domain.Category,
	difficulty domain.Difficulty,
	inputer domain.GameInputer,
	outputer domain.GameOutputer,
) (err error) {
	word, err := ChoiceWord(category, difficulty)
	if err != nil {
		return fmt.Errorf("choice word: %w", err)
	}

	game := domain.NewGame(word)

	for !game.IsFinished() {
		outputer.ShowGame(game)

		letter := inputer.GetLetter()

		game.Guess(letter)
	}

	outputer.ShowGameResult(game)

	return nil
}
