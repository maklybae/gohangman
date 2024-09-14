package infrastructure

import (
	"flag"
	"fmt"
	"makly/hangman/internal/domain"
	"makly/hangman/pkg/climenu"
)

func InitFlagsParameters() (path string, difficulty domain.Difficulty) {
	flag.StringVar(&path, "path", "", "path to json file with words collection")
	flag.Var(&difficulty, "difficulty", "difficulty level: easy, medium, hard")

	difficulty = domain.UnknownDifficulty
	flag.Parse()
	return
}

func ChooseDifficulty() (domain.Difficulty, error) {
	menu := climenu.NewMenu("Choose difficulty:")
	menu.AddItem(domain.EasyDifficulty.String())
	menu.AddItem(domain.MediumDifficulty.String())
	menu.AddItem(domain.HardDifficulty.String())
	chosenIndex, err := menu.RunMenu()
	if err != nil {
		return domain.UnknownDifficulty, fmt.Errorf("choose difficulty: %w", err)
	}
	return domain.Difficulty(chosenIndex), nil
}

func ChooseCategory(categories []domain.Category) (*domain.Category, error) {
	menu := climenu.NewMenu("Choose category:")
	for _, category := range categories {
		menu.AddItem(category.Name)
	}
	chosenIndex, err := menu.RunMenu()
	if err != nil {
		return nil, fmt.Errorf("choose category: %w", err)
	}
	return &categories[chosenIndex], nil
}
