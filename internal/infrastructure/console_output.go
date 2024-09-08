package infrastructure

import (
	"fmt"
	"makly/hangman/internal/domain"
	"makly/hangman/internal/draw"
)

type ConsoleOutput struct {
}

func (c *ConsoleOutput) ShowState(state domain.State) {
	fmt.Print(draw.StringStates[state])
}

func (c *ConsoleOutput) ShowGame(game *domain.Game) {
	state := game.State()
	c.ShowState(state)
	fmt.Printf("\n\n%s\n", game.Pattern())
}

func (c *ConsoleOutput) ShowGameResult(game *domain.Game) {
	if game.IsWin() {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost!")
	}
}
