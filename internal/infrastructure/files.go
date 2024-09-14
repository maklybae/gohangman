package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"makly/hangman/internal/domain"
	"os"
)

func ReadCollectionFromFile(path string) (wordsCollection *domain.WordsCollection, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			if err != nil {
				err = errors.Join(err, closeErr)
				return
			}

			err = fmt.Errorf("close file: %w", closeErr)
		}
	}()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	err = json.Unmarshal(byteValue, &wordsCollection)
	if err != nil {
		return nil, fmt.Errorf("unmarshal json: %w", err)
	}

	return wordsCollection, nil
}
