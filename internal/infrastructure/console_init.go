package infrastructure

import (
	"flag"
	"fmt"
	"log/slog"
	"path/filepath"

	"makly/hangman/internal/application"
	"makly/hangman/internal/domain"
	"makly/hangman/pkg/climenu"
)

func InitFlagsParameters() (path string, difficulty domain.Difficulty, maxMistakes int) {
	flag.StringVar(&path, "path", "", "path to json file with words collection")
	flag.Var(&difficulty, "difficulty", "difficulty level: easy, medium, hard")
	flag.IntVar(&maxMistakes, "maxmistakes", domain.StateCount, "maximum number of mistakes: integer from 1 to 26; default value: 6")
	difficulty = domain.UnknownDifficulty

	flag.Parse()

	return path, difficulty, maxMistakes
}

func ChooseDifficulty(menu climenu.MenuProvider) (difficulty domain.Difficulty, err error) {
	menu.AddItem("Secret difficulty (difficulty will be chosen randomly)")
	menu.AddItem(domain.EasyDifficulty.String())
	menu.AddItem(domain.MediumDifficulty.String())
	menu.AddItem(domain.HardDifficulty.String())

	slog.Info("Start choose difficulty menu", slog.Any("menu", menu))

	chosenIndex, err := menu.RunMenu()
	if err != nil {
		return domain.UnknownDifficulty, fmt.Errorf("choose difficulty: %w", err)
	}

	if chosenIndex == 0 {
		difficulty, err := application.ChoiceDifficulty()
		if err != nil {
			return domain.UnknownDifficulty, fmt.Errorf("random choose difficulty: %w", err)
		}

		slog.Info("Random chosen difficulty", slog.String("difficulty", difficulty.String()))

		return difficulty, nil
	}

	slog.Info("Chosen difficulty", slog.String("difficulty", domain.Difficulty(chosenIndex).String()))

	return domain.Difficulty(chosenIndex - 1), nil
}

func ChooseCategory(categories []domain.Category, menu climenu.MenuProvider) (category *domain.Category, err error) {
	menu.AddItem("Secret category (category will be chosen randomly)")

	for _, category := range categories {
		menu.AddItem(category.Name)
	}

	slog.Info("Start choose category menu", slog.Any("menu", menu))

	chosenIndex, err := menu.RunMenu()
	if err != nil {
		return nil, fmt.Errorf("choose category: %w", err)
	}

	if chosenIndex == 0 {
		category, err = application.ChoiceCategory(categories)
		if err != nil {
			return nil, fmt.Errorf("random choose category: %w", err)
		}

		slog.Info("Random chosen category", slog.String("category", category.Name))

		return category, nil
	}

	slog.Info("Chosen category", slog.String("category", categories[chosenIndex-1].Name))

	return &categories[chosenIndex-1], nil
}

func Init(defaultSamplePath, schemaPath string) (category *domain.Category, difficulty domain.Difficulty, maxMistakes int, err error) {
	jsonAbsPath, difficulty, maxMistakes := InitFlagsParameters()
	if jsonAbsPath == "" {
		jsonAbsPath, err = filepath.Abs(defaultSamplePath)
		if err != nil {
			return nil, domain.UnknownDifficulty, -1, fmt.Errorf("get absolute path: %w", err)
		}
	}

	slog.Info("Flags parsed",
		slog.String("path", jsonAbsPath),
		slog.String("difficulty", difficulty.String()),
		slog.Int("maxMistakes", maxMistakes))

	if maxMistakes < 1 || maxMistakes > ('Z'-'A'+1) {
		slog.Warn("Invalid maxMistakes value, set default value", slog.Int("maxMistakes", maxMistakes))
		maxMistakes = domain.StateCount
	}

	wordsCollection, err := ReadCollectionFromFile(jsonAbsPath, schemaPath)
	if err != nil {
		return nil, domain.UnknownDifficulty, -1, fmt.Errorf("read collection from file: %w", err)
	} else if wordsCollection == nil || len(wordsCollection.Categories) == 0 {
		return nil, domain.UnknownDifficulty, -1, &domain.BadWordsCollectionError{Message: "words collection is empty"}
	}

	slog.Info("Read words collection", slog.Any("words collection", wordsCollection))

	if difficulty == domain.UnknownDifficulty {
		difficulty, err = ChooseDifficulty(climenu.NewMenu("Choose difficulty:"))
		if err != nil {
			return nil, domain.UnknownDifficulty, -1, fmt.Errorf("start choose difficulty menu: %w", err)
		}
	}

	category, err = ChooseCategory(wordsCollection.Categories, climenu.NewMenu("Choose category:"))
	if err != nil {
		return nil, domain.UnknownDifficulty, -1, fmt.Errorf("choose category: %w", err)
	} else if category == nil || len(category.EasyWords)+len(category.MediumWords)+len(category.HardWords) == 0 {
		return nil, domain.UnknownDifficulty, -1, &domain.BadCategoryError{Message: "category is empty"}
	}

	return category, difficulty, maxMistakes, nil
}
