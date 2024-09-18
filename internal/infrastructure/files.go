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

type Reader interface {
	io.Reader
}

func ReadCollection(jsonReader, schemaReader Reader) (wordsCollection *domain.WordsCollection, err error) {
	jsonBytes, err := io.ReadAll(jsonReader)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	schemaBytes, err := io.ReadAll(schemaReader)
	if err != nil {
		return nil, fmt.Errorf("read schema file: %w", err)
	}

	documentLoader := gojsonschema.NewBytesLoader(jsonBytes)
	schemaLoader := gojsonschema.NewBytesLoader(schemaBytes)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return nil, fmt.Errorf("json schema validation: %w", err)
	}

	slog.Info("Validate json schema", slog.Bool("bool result", result.Valid()))

	if !result.Valid() {
		return nil, &IncorrectJSONError{Message: "json schema is invalid", JSONErrors: result.Errors()}
	}

	err = json.Unmarshal(jsonBytes, &wordsCollection)
	if err != nil {
		return nil, fmt.Errorf("unmarshal json: %w", err)
	}

	slog.Info("Unmarshal json", slog.Any("words collection", wordsCollection))

	return wordsCollection, nil
}

func ReadCollectionFromFile(jsonPath, schemaPath string) (wordsCollection *domain.WordsCollection, err error) {
	slog.Info("Open json file", slog.String("path", jsonPath))

	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}

	defer func() {
		if closeErr := jsonFile.Close(); closeErr != nil {
			if err != nil {
				err = errors.Join(err, closeErr)
				return
			}

			err = fmt.Errorf("close file: %w", closeErr)
		}

		slog.Info("Close json file", slog.String("path", jsonPath))
	}()

	slog.Info("Open json schema file", slog.String("path", schemaPath))

	schemaFile, err := os.Open(schemaPath)
	if err != nil {
		return nil, fmt.Errorf("open schema file: %w", err)
	}

	defer func() {
		if closeErr := schemaFile.Close(); closeErr != nil {
			if err != nil {
				err = errors.Join(err, closeErr)
				return
			}

			err = fmt.Errorf("close schema file: %w", closeErr)
		}

		slog.Info("Close json schema file", slog.String("path", schemaPath))
	}()

	return ReadCollection(jsonFile, schemaFile)
}
