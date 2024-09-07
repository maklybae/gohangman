package domain

type Game struct {
	attempts int
	mistakes int
	word     string
	guessed  map[rune]bool
	used     map[rune]bool
}

func NewGame(word string) *Game {
	return &Game{
		attempts: 0,
		mistakes: 0,
		word:     word,
		guessed:  make(map[rune]bool),
		used:     make(map[rune]bool),
	}
}

func (g *Game) GetAttempts() int {
	return g.attempts
}

func (g *Game) GetMistakes() int {
	return g.mistakes
}

func (g *Game) GetPattern() string {
	pattern := ""
	for _, letter := range g.word {
		if g.guessed[letter] {
			pattern += string(letter)
		} else {
			pattern += "_"
		}
	}
	return pattern
}

func (g *Game) Guess(letter rune) bool {
	return false
}

func (g *Game) IsFinished() bool {
	return false
}
