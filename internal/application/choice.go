package application

import (
	"crypto/rand"
	"fmt"
	"makly/hangman/internal/domain"
	"math/big"
)

type WordRandomizer interface {
	ChoiceWord(category *domain.Category, difficulty domain.Difficulty) (word *domain.Word, err error)
}

type RandomDefault struct{}

func ChoiceDifficulty() (difficulty domain.Difficulty, err error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(domain.DifficultyCount)))
	if err != nil {
		return -1, fmt.Errorf("random choose word: %w", err)
	}

	return domain.Difficulty(n.Int64()), nil
}

func ChoiceCategory(categories []domain.Category) (category *domain.Category, err error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(categories))))
	if err != nil {
		return nil, fmt.Errorf("random choose category: %w", err)
	}

	return &categories[n.Int64()], nil
}

func (rd *RandomDefault) ChoiceWord(category *domain.Category, difficulty domain.Difficulty) (word *domain.Word, err error) {
	var words []domain.Word

	switch difficulty {
	case domain.EasyDifficulty:
		words = category.EasyWords
	case domain.MediumDifficulty:
		words = category.MediumWords
	case domain.HardDifficulty:
		words = category.HardWords
	case domain.UnknownDifficulty:
		return nil, &domain.BadCategoryError{Message: "unknown difficulty"}
	default:
		return nil, &domain.BadCategoryError{Message: "unknown difficulty"}
	}

	if len(words) == 0 {
		return nil, &domain.BadCategoryError{Message: "words list for chosen category and difficulty is empty"}
	}

	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(words))))
	if err != nil {
		return nil, fmt.Errorf("random choose word: %w", err)
	}

	return &words[n.Int64()], nil
}
