package domain

type Game struct {
	attempts       int
	mistakes       int
	word           string
	correctLetters map[rune]bool
	used           map[rune]bool
}

func NewGame(word string) *Game {
	correctLetters := make(map[rune]bool)
	used := make(map[rune]bool)
	for _, letter := range word {
		correctLetters[letter] = true
		if letter == ' ' {
			used[letter] = true
		}
	}

	return &Game{
		attempts:       0,
		mistakes:       0,
		word:           word,
		correctLetters: correctLetters,
		used:           used,
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
		if g.used[letter] && g.correctLetters[letter] {
			pattern += string(letter)
		} else {
			pattern += "_"
		}
	}
	return pattern
}

func (g *Game) Guess(letter rune) {
	if g.used[letter] || letter == ' ' {
		return
	}

	g.attempts++
	g.used[letter] = true
	if g.correctLetters[letter] {
		return
	}

	g.mistakes++
}

func (g *Game) IsWin() bool {
	for _, letter := range g.word {
		if !g.used[letter] {
			return false
		}
	}
	return true
}

func (g *Game) IsLose() bool {
	return g.mistakes >= 6
}

func (g *Game) IsFinished() bool {
	return g.IsLose() || g.IsWin()
}
