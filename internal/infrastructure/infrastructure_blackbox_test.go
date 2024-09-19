package infrastructure_test

import (
	"bytes"
	"encoding/json"
	"makly/hangman/internal/domain"
	"makly/hangman/internal/infrastructure"
	"makly/hangman/internal/infrastructure/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const testStringSchema = `{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "type": "object",
    "properties": {
        "creator": {
            "type": "string"
        },
        "description": {
            "type": "string"
        },
        "categories": {
            "type": "array",
            "minItems": 1,
            "items": {
                "type": "object",
                "properties": {
                    "name": {
                        "type": "string"
                    },
                    "easy": {
                        "type": "array",
                        "items": {
                            "type": "object",
                            "properties": {
                                "word": {
                                    "type": "string",
                                    "pattern": "^[A-Za-z ]+$"
                                },
                                "hint": {
                                    "type": "string"
                                }
                            },
                            "required": [
                                "word",
                                "hint"
                            ]
                        }
                    },
                    "medium": {
                        "type": "array",
                        "items": {
                            "type": "object",
                            "properties": {
                                "word": {
                                    "type": "string",
                                    "pattern": "^[A-Za-z ]+$"
                                },
                                "hint": {
                                    "type": "string"
                                }
                            },
                            "required": [
                                "word",
                                "hint"
                            ]
                        }
                    },
                    "hard": {
                        "type": "array",
                        "items": {
                            "type": "object",
                            "properties": {
                                "word": {
                                    "type": "string",
                                    "pattern": "^[A-Za-z ]+$"
                                },
                                "hint": {
                                    "type": "string"
                                }
                            },
                            "required": [
                                "word",
                                "hint"
                            ]
                        }
                    }
                },
                "required": [
                    "name",
                    "easy",
                    "medium",
                    "hard"
                ]
            }
        }
    },
    "required": [
        "creator",
        "description",
        "categories"
    ]
}`

func TestValidateJSONBytes(t *testing.T) {
	tests := []struct {
		name        string
		jsonBytes   []byte
		expectError bool
	}{
		{
			name: "valid JSON",
			jsonBytes: []byte(`{
                "creator": "John Doe",
                "description": "Sample description",
                "categories": [
                    {
                        "name": "Category1",
                        "easy": [{"word": "apple", "hint": "A fruit"}],
                        "medium": [{"word": "banana", "hint": "Another fruit"}],
                        "hard": [{"word": "cherry", "hint": "Yet another fruit"}]
                    }
                ]
            }`),
			expectError: false,
		},
		{
			name: "valid JSON - multiple categories",
			jsonBytes: []byte(`{
                "creator": "Jane Doe",
                "description": "Another sample description",
                "categories": [
                    {
                        "name": "Category1",
                        "easy": [{"word": "apple", "hint": "A fruit"}],
                        "medium": [{"word": "banana", "hint": "Another fruit"}],
                        "hard": [{"word": "cherry", "hint": "Yet another fruit"}]
                    },
                    {
                        "name": "Category2",
                        "easy": [{"word": "dog", "hint": "An animal"}],
                        "medium": [{"word": "cat", "hint": "Another animal"}],
                        "hard": [{"word": "elephant", "hint": "A large animal"}]
                    }
                ]
            }`),
			expectError: false,
		},
		{
			name: "invalid JSON - missing hard category",
			jsonBytes: []byte(`{
                "creator": "John Doe",
                "description": "Sample description",
                "categories": [
                    {
                        "name": "Category1",
                        "easy": [{"word": "apple", "hint": "A fruit"}],
                        "medium": [{"word": "banana", "hint": "Another fruit"}]
                    }
                ]
            }`),
			expectError: true,
		},
		{
			name: "invalid JSON - missing creator",
			jsonBytes: []byte(`{
                "description": "Sample description",
                "categories": [
                    {
                        "name": "Category1",
                        "easy": [{"word": "apple", "hint": "A fruit"}],
                        "medium": [{"word": "banana", "hint": "Another fruit"}],
                        "hard": [{"word": "cherry", "hint": "Yet another fruit"}]
                    }
                ]
            }`),
			expectError: true,
		},
		{
			name: "invalid JSON - missing description",
			jsonBytes: []byte(`{
                "creator": "John Doe",
                "categories": [
                    {
                        "name": "Category1",
                        "easy": [{"word": "apple", "hint": "A fruit"}],
                        "medium": [{"word": "banana", "hint": "Another fruit"}],
                        "hard": [{"word": "cherry", "hint": "Yet another fruit"}]
                    }
                ]
            }`),
			expectError: true,
		},
		{
			name: "invalid JSON - missing categories",
			jsonBytes: []byte(`{
                "creator": "John Doe",
                "description": "Sample description"
            }`),
			expectError: true,
		},
		{
			name: "invalid JSON - empty categories",
			jsonBytes: []byte(`{
                "creator": "John Doe",
                "description": "Sample description",
                "categories": []
            }`),
			expectError: true,
		},
		{
			name: "invalid JSON - invalid word pattern",
			jsonBytes: []byte(`{
                "creator": "John Doe",
                "description": "Sample description",
                "categories": [
                    {
                        "name": "Category1",
                        "easy": [{"word": "apple123", "hint": "A fruit"}],
                        "medium": [{"word": "banana", "hint": "Another fruit"}],
                        "hard": [{"word": "cherry", "hint": "Yet another fruit"}]
                    }
                ]
            }`),
			expectError: true,
		},
	}

	var jsonErr *infrastructure.IncorrectJSONError

	assertInstance := assert.New(t)
	validator := &infrastructure.Validator{}

	for _, tt := range tests {
		schemaReader := bytes.NewReader([]byte(testStringSchema))
		err := validator.ValidateJSONBytes(tt.jsonBytes, schemaReader)

		if tt.expectError {
			assertInstance.ErrorAs(err, &jsonErr, tt.name)
		} else {
			assertInstance.NoError(err, tt.name)
		}
	}
}

func TestReadCollection(t *testing.T) {
	mockJSONBytesValidator := &mocks.JSONBytesValidator{}
	mockJSONBytesValidator.On("ValidateJSONBytes", mock.Anything, mock.Anything).Return(nil)

	tests := []struct {
		name                    string
		jsonBytes               []byte
		expectedWordsCollection domain.WordsCollection
		expectedErr             error
	}{
		{
			name: "valid JSON",
			jsonBytes: []byte(`{
                "creator": "John Doe",
                "description": "Sample description",
                "categories": [
                    {
                        "name": "Category1",
                        "easy": [{"word": "apple", "hint": "A fruit"}],
                        "medium": [{"word": "banana", "hint": "Another fruit"}],
                        "hard": [{"word": "cherry", "hint": "Yet another fruit"}]
                    }
                ]
            }`),
			expectedWordsCollection: domain.WordsCollection{
				Creator:     "John Doe",
				Description: "Sample description",
				Categories: []domain.Category{
					{
						Name: "Category1",
						EasyWords: []domain.Word{
							{Word: "apple", Hint: "A fruit"},
						},
						MediumWords: []domain.Word{
							{Word: "banana", Hint: "Another fruit"},
						},
						HardWords: []domain.Word{
							{Word: "cherry", Hint: "Yet another fruit"},
						},
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "invalid JSON - type mismatch",
			jsonBytes: []byte(`{
                "creator": "John Doe",
                "description": "Sample description",
                "categories": [
                    {
                        "name": "Category1",
                        "easy": [{"word": 123, "hint": "A fruit"}],
                        "medium": [{"word": "banana", "hint": "Another fruit"}],
                        "hard": [{"word": "cherry", "hint": "Yet another fruit"}]
                    }
                ]
            }`),
			expectedWordsCollection: domain.WordsCollection{},
			expectedErr:             &json.UnmarshalTypeError{},
		},
		{
			name: "invalid JSON - unexpected end of JSON input",
			jsonBytes: []byte(`{
            "creator": "John Doe",
            "description": "Sample description",
            "categories": [
                {
                "name": "Category1",
                "easy": [{"word": "apple", "hint": "A fruit"}],
                "medium": [{"word": "banana", "hint": "Another fruit"}],
                "hard": [{"word": "cherry", "hint": "Yet another fruit"}]
                }
            ]`), // Missing closing brace for the object
			expectedWordsCollection: domain.WordsCollection{},
			expectedErr:             &json.SyntaxError{},
		},
		{
			name: "invalid JSON - syntax error",
			jsonBytes: []byte(`{
                "creator": "John Doe",
                "description": "Sample description",
                "categories": [
                    {
                        "name": "Category1",
                        "easy": [{"word": "apple", "hint": "A fruit"}],
                        "medium": [{"word": "banana", "hint": "Another fruit"}],
                        "hard": [{"word": "cherry", "hint": "Yet another fruit"}]
                    }
                ]`), // Another missing closing brace for the object
			expectedWordsCollection: domain.WordsCollection{},
			expectedErr:             &json.SyntaxError{},
		},
	}

	assertInstance := assert.New(t)

	for _, tt := range tests {
		wordsCollection, err := infrastructure.ReadCollection(
			bytes.NewReader(tt.jsonBytes), nil, mockJSONBytesValidator, // nil is used for schemaReader
		)

		if tt.expectedErr != nil {
			assertInstance.ErrorAs(err, &tt.expectedErr, tt.name)
		} else {
			assertInstance.Equal(tt.expectedWordsCollection, *wordsCollection, tt.name)
		}
	}
}
