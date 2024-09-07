package application

import (
	"makly/hangman/internal/domain"
	"math/rand"
)

func ChoiceWord(category *domain.Category, difficulty domain.Difficulty) string {
	var words []string
	switch difficulty {
	case domain.EasyDifficulty:
		words = category.EasyWords
	case domain.MediumDifficulty:
		words = category.MediumWords
	case domain.HardDifficulty:
		words = category.HardWords
	}
	return words[rand.Intn(len(words))]
}

func ChoiceCategory(wordsCollection *domain.WordsCollection) *domain.Category {
	return &wordsCollection.Categories[rand.Intn(len(wordsCollection.Categories))]
}
