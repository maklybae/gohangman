package application

import (
	"errors"
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
	reshow := true

	for !game.IsFinished() {
		if reshow {
			outputer.ShowGame(game)
		}

		letter, err := inputer.GetLetter()
		if err != nil {
			var inputerError *domain.InputerError
			if errors.As(err, &inputerError) {
				reshow = false

				outputer.ShowInputError(err)

				continue
			}

			return fmt.Errorf("getting letter to guess: %w", err)
		}

		game.Guess(letter)

		reshow = true
	}

	outputer.ShowGame(game)
	outputer.ShowGameResult(game)

	return nil
}
