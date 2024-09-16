package infrastructure

import (
	"flag"
	"fmt"
	"log/slog"
	"makly/hangman/internal/domain"
	"makly/hangman/pkg/climenu"
	"path/filepath"
)

func InitFlagsParameters() (path string, difficulty domain.Difficulty) {
	flag.StringVar(&path, "path", "", "path to json file with words collection")
	flag.Var(&difficulty, "difficulty", "difficulty level: easy, medium, hard")
	difficulty = domain.UnknownDifficulty

	flag.Parse()

	return path, difficulty
}

func ChooseDifficulty() (domain.Difficulty, error) {
	menu := climenu.NewMenu("Choose difficulty:")
	menu.AddItem(domain.EasyDifficulty.String())
	menu.AddItem(domain.MediumDifficulty.String())
	menu.AddItem(domain.HardDifficulty.String())

	slog.Info("Start choose difficulty menu", slog.Any("menu", menu))

	chosenIndex, err := menu.RunMenu()
	if err != nil {
		return domain.UnknownDifficulty, fmt.Errorf("choose difficulty: %w", err)
	}

	slog.Info("Chosen difficulty", slog.String("difficulty", domain.Difficulty(chosenIndex).String()))

	return domain.Difficulty(chosenIndex), nil
}

func ChooseCategory(categories []domain.Category) (*domain.Category, error) {
	menu := climenu.NewMenu("Choose category:")
	for _, category := range categories {
		menu.AddItem(category.Name)
	}

	slog.Info("Start choose category menu", slog.Any("menu", menu))
	chosenIndex, err := menu.RunMenu()

	if err != nil {
		return nil, fmt.Errorf("choose category: %w", err)
	}

	slog.Info("Chosen category", slog.String("category", categories[chosenIndex].Name))

	return &categories[chosenIndex], nil
}

func Init(defaultSamplePath, schemaPath string) (category *domain.Category, difficulty domain.Difficulty, err error) {
	jsonAbsPath, difficulty := InitFlagsParameters()
	if jsonAbsPath == "" {
		jsonAbsPath, err = filepath.Abs(defaultSamplePath)
		if err != nil {
			return nil, domain.UnknownDifficulty, fmt.Errorf("get absolute path: %w", err)
		}
	}

	slog.Info("Flags parsed", slog.String("path", jsonAbsPath), slog.String("difficulty", difficulty.String()))

	wordsCollection, err := ReadCollectionFromFile(jsonAbsPath, schemaPath)
	if err != nil {
		return nil, domain.UnknownDifficulty, fmt.Errorf("read collection from file: %w", err)
	} else if wordsCollection == nil || len(wordsCollection.Categories) == 0 {
		return nil, domain.UnknownDifficulty, &domain.BadWordsCollectionError{Message: "words collection is empty"}
	}

	slog.Info("Read words collection", slog.Any("words collection", wordsCollection))

	if difficulty == domain.UnknownDifficulty {
		difficulty, err = ChooseDifficulty()
		if err != nil {
			return nil, domain.UnknownDifficulty, fmt.Errorf("start choose difficulty menu: %w", err)
		}
	}

	category, err = ChooseCategory(wordsCollection.Categories)
	if err != nil {
		return nil, domain.UnknownDifficulty, fmt.Errorf("choose category: %w", err)
	} else if category == nil || len(category.EasyWords)+len(category.MediumWords)+len(category.HardWords) == 0 {
		return nil, domain.UnknownDifficulty, &domain.BadCategoryError{Message: "category is empty"}
	}

	return category, difficulty, nil
}
