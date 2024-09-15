package infrastructure

import (
	"fmt"
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

func (c *ConsoleOutput) showMistakes(mistakes int) {
	fmt.Printf("Mistakes: %d / %d\n", mistakes, domain.MaxMistakes)
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
	c.showMistakes(game.Mistakes())
	c.showUsed(game.Used())
	c.showState(game.State())
	fmt.Printf("\n\nPattern: %s\n\n", game.Pattern())
	fmt.Printf("Guess next letter: ")
}

func (c *ConsoleOutput) ShowGameResult(game *domain.Game) {
	if game.IsWin() {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost!")
	}
}

func (c *ConsoleOutput) ShowInputError(err error) {
	// Clear the last line with last input
	fmt.Printf("\033[1A")
	fmt.Printf("\033[2K")

	fmt.Printf("Game error: %s. Try again: ", err)
}
