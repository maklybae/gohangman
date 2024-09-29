package infrastructure

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/xeipuuv/gojsonschema"

	"makly/hangman/internal/domain"
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

type JSONBytesValidator interface {
	ValidateJSONBytes(jsonBytes []byte, schemaReader Reader) (err error)
}

type Validator struct{}

func (v *Validator) ValidateJSONBytes(jsonBytes []byte, schemaReader Reader) (err error) {
	schemaBytes, err := io.ReadAll(schemaReader)
	if err != nil {
		return fmt.Errorf("read schema file: %w", err)
	}

	documentLoader := gojsonschema.NewBytesLoader(jsonBytes)
	schemaLoader := gojsonschema.NewBytesLoader(schemaBytes)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return fmt.Errorf("json schema validation: %w", err)
	}

	slog.Info("Validate json schema", slog.Bool("bool result", result.Valid()))

	if !result.Valid() {
		return &IncorrectJSONError{Message: "json schema is invalid", JSONErrors: result.Errors()}
	}

	return nil
}

func ReadCollection(jsonReader, schemaReader Reader, validator JSONBytesValidator) (wordsCollection *domain.WordsCollection, err error) {
	jsonBytes, err := io.ReadAll(jsonReader)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	if err := validator.ValidateJSONBytes(jsonBytes, schemaReader); err != nil {
		return nil, fmt.Errorf("reading collection: %w", err)
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

	return ReadCollection(jsonFile, schemaFile, &Validator{})
}
