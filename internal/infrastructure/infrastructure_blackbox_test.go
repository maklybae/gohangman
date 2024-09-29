package infrastructure_test

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"makly/hangman/internal/domain"
	"makly/hangman/internal/infrastructure"
	"makly/hangman/internal/infrastructure/mocks"
	menuMocks "makly/hangman/pkg/climenu/mocks"
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
	log.SetOutput(io.Discard)

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
	log.SetOutput(io.Discard)

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

func TestInitFlagsParameters(t *testing.T) {
	tests := []struct {
		name                string
		args                []string
		expectedPath        string
		expectedDifficulty  domain.Difficulty
		expectedMaxMistakes int
	}{
		{
			name:                "default values",
			args:                []string{},
			expectedPath:        "",
			expectedDifficulty:  domain.UnknownDifficulty,
			expectedMaxMistakes: domain.StateCount,
		},
		{
			name:                "valid arguments",
			args:                []string{"-path", "test/path", "-difficulty", "medium", "-maxmistakes", "5"},
			expectedPath:        "test/path",
			expectedDifficulty:  domain.MediumDifficulty,
			expectedMaxMistakes: 5,
		},
		{
			name:                "invalid difficulty",
			args:                []string{"-path", "test/path", "-difficulty", "invalid", "-maxmistakes", "5"},
			expectedPath:        "test/path",
			expectedDifficulty:  domain.UnknownDifficulty,
			expectedMaxMistakes: 5,
		},
		{
			name:                "missing max mistakes",
			args:                []string{"-path", "test/path", "-difficulty", "medium"},
			expectedPath:        "test/path",
			expectedDifficulty:  domain.MediumDifficulty,
			expectedMaxMistakes: domain.StateCount,
		},
		{
			name:                "no args",
			args:                []string{},
			expectedPath:        "",
			expectedDifficulty:  domain.UnknownDifficulty,
			expectedMaxMistakes: domain.StateCount,
		},
		{
			name:                "only path",
			args:                []string{"-path", "test/path"},
			expectedPath:        "test/path",
			expectedDifficulty:  domain.UnknownDifficulty,
			expectedMaxMistakes: domain.StateCount,
		},
	}

	for _, tt := range tests {
		// Reset flags and os.Args before each test
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

		os.Args = append([]string{"cmd"}, tt.args...)

		gotPath, gotDifficulty, gotMaxMistakes := infrastructure.InitFlagsParameters()

		assert.Equal(t, tt.expectedPath, gotPath)
		assert.Equal(t, tt.expectedDifficulty, gotDifficulty)
		assert.Equal(t, tt.expectedMaxMistakes, gotMaxMistakes)
	}
}

func TestChooseDifficulty(t *testing.T) {
	log.SetOutput(io.Discard)

	assertInstance := assert.New(t)
	mockMenu := &menuMocks.MenuProvider{}
	mockMenu.On("AddItem", mock.Anything).Return()

	mockMenu.On("RunMenu").Return(1, nil).Once()
	difficulty, err := infrastructure.ChooseDifficulty(mockMenu)
	assertInstance.NoError(err)
	assertInstance.Equal(domain.EasyDifficulty, difficulty)

	mockMenu.On("RunMenu").Return(2, nil).Once()
	difficulty, err = infrastructure.ChooseDifficulty(mockMenu)
	assertInstance.NoError(err)
	assertInstance.Equal(domain.MediumDifficulty, difficulty)

	mockMenu.On("RunMenu").Return(3, nil).Once()
	difficulty, err = infrastructure.ChooseDifficulty(mockMenu)
	assertInstance.NoError(err)
	assertInstance.Equal(domain.HardDifficulty, difficulty)

	mockMenu.On("RunMenu").Return(0, nil).Once()
	difficulty, err = infrastructure.ChooseDifficulty(mockMenu)
	assertInstance.NoError(err)
	assertInstance.Contains([]domain.Difficulty{domain.EasyDifficulty, domain.MediumDifficulty, domain.HardDifficulty}, difficulty)
}

func TestChooseCategory(t *testing.T) {
	log.SetOutput(io.Discard)

	categories := []domain.Category{
		{
			Name: "Category1",
		},
		{
			Name: "Category2",
		},
		{
			Name: "Category3",
		},
	}

	assertInstance := assert.New(t)
	mockMenu := &menuMocks.MenuProvider{}
	mockMenu.On("AddItem", mock.Anything).Return()

	mockMenu.On("RunMenu").Return(1, nil).Once()
	category, err := infrastructure.ChooseCategory(categories, mockMenu)
	assertInstance.NoError(err)
	assertInstance.Equal("Category1", category.Name)

	mockMenu.On("RunMenu").Return(2, nil).Once()
	category, err = infrastructure.ChooseCategory(categories, mockMenu)
	assertInstance.NoError(err)
	assertInstance.Equal("Category2", category.Name)

	mockMenu.On("RunMenu").Return(3, nil).Once()
	category, err = infrastructure.ChooseCategory(categories, mockMenu)
	assertInstance.NoError(err)
	assertInstance.Equal("Category3", category.Name)

	mockMenu.On("RunMenu").Return(0, nil).Once()
	category, err = infrastructure.ChooseCategory(categories, mockMenu)
	assertInstance.NoError(err)
	assertInstance.Contains([]string{"Category1", "Category2", "Category3"}, category.Name)
}
