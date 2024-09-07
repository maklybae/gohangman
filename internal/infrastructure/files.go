package infrastructure

import (
	"encoding/json"
	"io"
	"makly/hangman/internal/domain"
	"os"
)

func ReadCollectionFromFile(path string) (*domain.WordsCollection, error) {
	wordsCollection := domain.WordsCollection{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteValue, &wordsCollection)
	if err != nil {
		return nil, err
	}

	return &wordsCollection, nil
}
