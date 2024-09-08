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
