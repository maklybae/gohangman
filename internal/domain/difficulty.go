package domain

import "fmt"

type Difficulty int

const (
	EasyDifficulty Difficulty = iota
	MediumDifficulty
	HardDifficulty
	UnknownDifficulty
)

const DifficultyCount = 3

func (d Difficulty) String() string {
	return [...]string{"Easy", "Medium", "Hard", "Unknown"}[d]
}

func (d *Difficulty) Set(value string) error {
	switch value {
	case "easy":
		*d = EasyDifficulty
	case "medium":
		*d = MediumDifficulty
	case "hard":
		*d = HardDifficulty
	default:
		*d = UnknownDifficulty
	}

	return nil
}

type BadDifficultyError struct {
	Message string
}

func (e *BadDifficultyError) Error() string {
	return fmt.Sprintf("bad difficulty: %s", e.Message)
}
