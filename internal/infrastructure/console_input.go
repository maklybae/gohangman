package infrastructure

import (
	"bufio"
	"fmt"
	"log/slog"
	"makly/hangman/internal/domain"
	"os"
	"unicode"
)

type ConsoleInput struct {
	scanner *bufio.Scanner
}

func NewConsoleInput() *ConsoleInput {
	return &ConsoleInput{scanner: bufio.NewScanner(os.Stdin)}
}

func (c *ConsoleInput) GetLetter() (letter rune, err error) {
	c.scanner.Scan()

	err = c.scanner.Err()
	if err != nil {
		return 0, fmt.Errorf("getting letter via bufio: %w", err)
	}

	text := c.scanner.Text()

	slog.Info("Got letter from standard cin", slog.String("letter", text))

	if len([]rune(text)) != 1 {
		return 0, &domain.InputerError{Message: "not a single letter", InnerError: nil}
	}

	letter = rune(text[0])
	if (letter < 'a' || letter > 'z') && (letter < 'A' || letter > 'Z') {
		return 0, &domain.InputerError{Message: "letter validation", InnerError: nil}
	}

	return unicode.ToLower(letter), nil
}
