package domain

import "log/slog"

const MaxMistakes = 6

type Game struct {
	attempts       int
	mistakes       int
	word           Word
	correctLetters map[rune]bool
	used           map[rune]bool
}

func NewGame(word *Word) *Game {
	correctLetters := make(map[rune]bool)
	used := make(map[rune]bool)

	for _, letter := range word.Word {
		correctLetters[letter] = true

		if letter == ' ' {
			used[letter] = true
		}
	}

	return &Game{
		attempts:       0,
		mistakes:       0,
		word:           *word,
		correctLetters: correctLetters,
		used:           used,
	}
}

func (g *Game) Attempts() int {
	return g.attempts
}

func (g *Game) Mistakes() int {
	return g.mistakes
}

func (g *Game) State() State {
	switch g.mistakes {
	case 1:
		return Head
	case 2:
		return Body
	case 3:
		return LeftArm
	case 4:
		return RightArm
	case 5:
		return LeftLeg
	case 6:
		return RightLeg
	default:
		return Initial
	}
}

func (g *Game) Used() map[rune]bool {
	return g.used
}

func (g *Game) Pattern() string {
	pattern := ""

	for _, letter := range g.word.Word {
		if g.used[letter] && g.correctLetters[letter] {
			pattern += string(letter)
		} else {
			pattern += "_"
		}
	}

	slog.Info("Pattern generated", slog.String("pattern", pattern), slog.Any("game", g))

	return pattern
}

func (g *Game) Hint() string {
	return g.word.Hint
}

func (g *Game) Guess(letter rune) {
	slog.Info("Guess letter", slog.String("letter", string(letter)))

	if g.used[letter] || letter == ' ' {
		return
	}

	g.attempts++
	g.used[letter] = true

	if g.correctLetters[letter] {
		slog.Info("Correct guess", slog.String("letter", string(letter)))
		return
	}

	slog.Info("Incorrect guess", slog.String("letter", string(letter)))

	g.mistakes++
}

func (g *Game) IsWin() bool {
	for _, letter := range g.word.Word {
		if !g.used[letter] {
			return false
		}
	}

	slog.Info("Win check completed successfully", slog.Any("game", g))

	return true
}

func (g *Game) IsLose() bool {
	return g.mistakes >= 6
}

func (g *Game) IsFinished() bool {
	return g.IsLose() || g.IsWin()
}

func (g *Game) IsHintAvailable() bool {
	return g.mistakes >= 3
}

func (g *Game) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int("attempts", g.attempts),
		slog.Int("mistakes", g.mistakes),
		slog.String("word", g.word.Word),
	)
}
