package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"makly/hangman/internal/domain"
	"os"

	"github.com/xeipuuv/gojsonschema"
)

type IncorrectJSONError struct {
	Message    string
	JSONErrors []gojsonschema.ResultError
}

func (e *IncorrectJSONError) Error() string {
	errorString := fmt.Sprintf("incorrect json: %s\n", e.Message)
	for _, desc := range e.JSONErrors {
		errorString += fmt.Sprintf("- %s\n", desc)
	}

	return errorString
}

func ReadCollectionFromFile(jsonPath, schemaPath string) (wordsCollection *domain.WordsCollection, err error) {
	file, err := os.Open(jsonPath)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}

	slog.Info("Open json file", slog.String("path", jsonPath))

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			if err != nil {
				err = errors.Join(err, closeErr)
				return
			}

			err = fmt.Errorf("close file: %w", closeErr)
		}

		slog.Info("Close json file", slog.String("path", jsonPath))
	}()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	slog.Info("Read json file", slog.String("path", jsonPath))

	schemaLoader := gojsonschema.NewReferenceLoader("file://" + schemaPath)
	documentLoader := gojsonschema.NewBytesLoader(byteValue)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return nil, fmt.Errorf("json schema validation: %w", err)
	}

	slog.Info("Validate json schema", slog.Bool("bool result", result.Valid()))

	if !result.Valid() {
		return nil, &IncorrectJSONError{Message: "json schema is invalid", JSONErrors: result.Errors()}
	}

	err = json.Unmarshal(byteValue, &wordsCollection)
	if err != nil {
		return nil, fmt.Errorf("unmarshal json: %w", err)
	}

	slog.Info("Unmarshal json", slog.Any("words collection", wordsCollection))

	return wordsCollection, nil
}
