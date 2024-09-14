package infrastructure

import (
	"fmt"
)

type ConsoleInput struct {
}

func NewConsoleInput() *ConsoleInput {
	return &ConsoleInput{}
}

func (c *ConsoleInput) GetLetter() rune {
	var letter rune
	fmt.Scanf("%c\n", &letter)

	return letter
}
