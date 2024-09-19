// package needs to test the GetLetter method without using scanner replacement (internal field of ConsoleInput)

package infrastructure //nolint

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLetter(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		returnValue rune
		expectError bool
	}{
		{
			name:        "valid letter",
			input:       "a",
			returnValue: 'a',
			expectError: false,
		},
		{
			name:        "valid letter - uppercase",
			input:       "A",
			returnValue: 'a',
			expectError: false,
		},
		{
			name:        "invalid letter - multiple characters",
			input:       "ab",
			returnValue: 0,
			expectError: true,
		},
		{
			name:        "invalid letter - non-alphabet character",
			input:       "1",
			returnValue: 0,
			expectError: true,
		},
		{
			name:        "valid letter - Russian lowercase",
			input:       "б",
			returnValue: 0,
			expectError: true,
		},
		{
			name:        "valid letter - Russian uppercase",
			input:       "Б",
			returnValue: 0,
			expectError: true,
		},
		{
			name:        "invalid letter - Russian multiple characters",
			input:       "аб",
			returnValue: 0,
			expectError: true,
		},
		{
			name:        "invalid letter - escape sequence",
			input:       "\n",
			returnValue: 0,
			expectError: true,
		},
		{
			name:        "invalid letter - tab character",
			input:       "\t",
			returnValue: 0,
			expectError: true,
		},
		{
			name:        "invalid letter - space character",
			input:       " ",
			returnValue: 0,
			expectError: true,
		},
		{
			name:        "valid letter - lowercase z",
			input:       "z",
			returnValue: 'z',
			expectError: false,
		},
		{
			name:        "valid letter - uppercase Z",
			input:       "Z",
			returnValue: 'z',
			expectError: false,
		},
		{
			name:        "valid letter - lowercase m",
			input:       "m",
			returnValue: 'm',
			expectError: false,
		},
		{
			name:        "valid letter - uppercase M",
			input:       "M",
			returnValue: 'm',
			expectError: false,
		},
		{
			name:        "valid letter - lowercase n",
			input:       "n",
			returnValue: 'n',
			expectError: false,
		},
		{
			name:        "valid letter - uppercase N",
			input:       "N",
			returnValue: 'n',
			expectError: false,
		},
		{
			name:        "valid letter - lowercase k",
			input:       "k",
			returnValue: 'k',
			expectError: false,
		},
		{
			name:        "valid letter - uppercase K",
			input:       "K",
			returnValue: 'k',
			expectError: false,
		},
	}

	assertInstance := assert.New(t)
	consoleInput := NewConsoleInput()

	for _, tt := range tests {
		consoleInput.scanner = bufio.NewScanner(bytes.NewReader([]byte(tt.input + "\n")))
		letter, err := consoleInput.GetLetter()

		if tt.expectError {
			assertInstance.Error(err)
		} else {
			assertInstance.NoError(err)
			assertInstance.Equal(tt.returnValue, letter)
		}
	}
}
