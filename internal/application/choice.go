package application

import (
	"crypto/rand"
	"fmt"
	"makly/hangman/internal/domain"
	"math/big"
)

func ChoiceWord(category *domain.Category, difficulty domain.Difficulty) (word *domain.Word, err error) {
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
