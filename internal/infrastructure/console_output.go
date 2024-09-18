package infrastructure

import (
	"fmt"
	"log/slog"
	"makly/hangman/internal/domain"
	"makly/hangman/internal/draw"
	"unicode"
)

type ConsoleOutput struct {
}

func NewConsoleOutput() *ConsoleOutput {
	return &ConsoleOutput{}
}

func (c *ConsoleOutput) showState(state domain.State) {
	fmt.Print(draw.StringStates[state])
}

func (c *ConsoleOutput) showUsed(used map[rune]bool) {
	fmt.Print("Used: ")

	for letter := 'a'; letter <= 'z'; letter++ {
		if used[letter] {
			// Bold the used letters
			fmt.Printf("\033[1m")
			fmt.Printf("%c ", unicode.ToUpper(letter))
			fmt.Printf("\033[0m")
		} else {
			fmt.Printf("%c ", letter)
		}
	}

	fmt.Println()
}

func (c *ConsoleOutput) showAttempts(attempts int) {
	fmt.Printf("Attempts: %d\n", attempts)
}

func (c *ConsoleOutput) showMistakes(mistakes, maxMistakes int) {
	fmt.Printf("Mistakes: %d / %d\n", mistakes, maxMistakes)
}

func (c *ConsoleOutput) showHint(hint string) {
	reversedHint := ""
	for _, r := range hint {
		reversedHint = string(r) + reversedHint
	}

	fmt.Printf("Reversed hint: %s\n", reversedHint)
}

func (c *ConsoleOutput) showPattern(pattern string) {
	fmt.Printf("Pattern: %s\n", pattern)
}

func (c *ConsoleOutput) clear() {
	fmt.Printf("\033[1A")
	fmt.Printf("\033[2K")
	fmt.Printf("\0338")
}

func (c *ConsoleOutput) memorizeCursor() {
	fmt.Printf("\0337")
}

func (c *ConsoleOutput) ShowGame(game *domain.Game) {
	if game.Attempts() != 0 {
		c.clear()
	}

	c.memorizeCursor()

	c.showAttempts(game.Attempts())
	c.showMistakes(game.Mistakes(), game.MaxMistakes())
	c.showUsed(game.Used())
	c.showState(game.State())
	fmt.Printf("\n\n")
	c.showPattern(game.Pattern())

	if game.IsHintAvailable() {
		c.showHint(game.Hint())
	}

	fmt.Printf("\n")
	fmt.Printf("Guess next letter: ")

	slog.Info("Current game state printed", slog.Any("game", game))
}

func (c *ConsoleOutput) ShowGameResult(game *domain.Game) {
	if game.IsWin() {
		fmt.Println("You won!")
		slog.Info("Game result printed", slog.String("result", "win"))
	} else {
		fmt.Println("You lost!")
		slog.Info("Game result printed", slog.String("result", "lose"))
	}
}

func (c *ConsoleOutput) ShowInputError(err error) {
	// Clear the last line with last input
	fmt.Printf("\033[1A")
	fmt.Printf("\033[2K")

	fmt.Printf("Game error: %s. Try again: ", err)
}
