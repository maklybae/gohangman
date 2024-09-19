package application

import (
	"errors"
	"fmt"
	"log/slog"
	"makly/hangman/internal/domain"
)

func RunGameSession(
	category *domain.Category,
	difficulty domain.Difficulty,
	maxMistakes int,
	inputer domain.GameInputer,
	outputer domain.GameOutputer,
	wordRandomizer WordRandomizer,
) (err error) {
	word, err := wordRandomizer.ChoiceWord(category, difficulty)
	if err != nil {
		return fmt.Errorf("choice word: %w", err)
	}

	slog.Info("Random choose word", slog.String("word", word.Word))

	game := domain.NewGame(word, maxMistakes)
	slog.Info("Game started", "game", game)

	reshow := true

	for !game.IsFinished() {
		if reshow {
			outputer.ShowGame(game)
			slog.Info("Reshow game", "game", game)
		}

		letter, err := inputer.GetLetter()
		if err != nil {
			var inputerError *domain.InputerError
			if errors.As(err, &inputerError) {
				slog.Error("Getting letter to guess", slog.Any("error", err))

				reshow = false

				outputer.ShowInputError(err)

				continue
			}

			return fmt.Errorf("getting letter to guess: %w", err)
		}

		slog.Info("Got correct letter to guess", slog.String("Letter", string(letter)))

		game.Guess(letter)

		reshow = true
	}

	outputer.ShowGame(game)
	outputer.ShowGameResult(game)

	return nil
}
