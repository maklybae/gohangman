package domain

type Difficulty int

const (
	EasyDifficulty Difficulty = iota
	MediumDifficulty
	HardDifficulty
	UnknownDifficulty
)

func (d Difficulty) String() string {
	return [...]string{"Easy", "Medium", "Hard"}[d]
}

func (d Difficulty) Set(value string) error {
	switch value {
	case "easy":
		d = EasyDifficulty
	case "medium":
		d = MediumDifficulty
	case "hard":
		d = HardDifficulty
	}
	return nil
}
