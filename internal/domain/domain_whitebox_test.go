package domain //nolint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameState(t *testing.T) {
	tests := []struct {
		name        string
		mistakes    int
		maxMistakes int
		expected    State
	}{
		{
			name:        "Initial state",
			mistakes:    0,
			maxMistakes: 6,
			expected:    Initial,
		},
		{
			name:        "Head state",
			mistakes:    1,
			maxMistakes: 6,
			expected:    Head,
		},
		{
			name:        "Body state",
			mistakes:    2,
			maxMistakes: 6,
			expected:    Body,
		},
		{
			name:        "LeftArm state",
			mistakes:    3,
			maxMistakes: 6,
			expected:    LeftArm,
		},
		{
			name:        "RightArm state",
			mistakes:    4,
			maxMistakes: 6,
			expected:    RightArm,
		},
		{
			name:        "LeftLeg state",
			mistakes:    5,
			maxMistakes: 6,
			expected:    LeftLeg,
		},
		{
			name:        "RightLeg state",
			mistakes:    6,
			maxMistakes: 6,
			expected:    RightLeg,
		},
		{
			name:        "Initial state with maxMistakes 4",
			mistakes:    0,
			maxMistakes: 4,
			expected:    Initial,
		},
		{
			name:        "Head state with maxMistakes 4",
			mistakes:    1,
			maxMistakes: 4,
			expected:    Head,
		},
		{
			name:        "LeftArm state with maxMistakes 4",
			mistakes:    2,
			maxMistakes: 4,
			expected:    LeftArm,
		},
		{
			name:        "RightArm state with maxMistakes 4",
			mistakes:    3,
			maxMistakes: 4,
			expected:    RightArm,
		},
		{
			name:        "RightLeg state with maxMistakes 4",
			mistakes:    4,
			maxMistakes: 4,
			expected:    RightLeg,
		},
		{
			name:        "Initial state with maxMistakes 8",
			mistakes:    0,
			maxMistakes: 8,
			expected:    Initial,
		},
		{
			name:        "Initial state with maxMistakes 8",
			mistakes:    1,
			maxMistakes: 8,
			expected:    Initial,
		},
		{
			name:        "Head state with maxMistakes 8",
			mistakes:    2,
			maxMistakes: 8,
			expected:    Head,
		},
		{
			name:        "Body state with maxMistakes 8",
			mistakes:    3,
			maxMistakes: 8,
			expected:    Body,
		},
		{
			name:        "LeftArm state with maxMistakes 8",
			mistakes:    4,
			maxMistakes: 8,
			expected:    LeftArm,
		},
		{
			name:        "LeftArm state with maxMistakes 8",
			mistakes:    5,
			maxMistakes: 8,
			expected:    LeftArm,
		},
		{
			name:        "RightArm state with maxMistakes 8",
			mistakes:    6,
			maxMistakes: 8,
			expected:    RightArm,
		},
		{
			name:        "LeftLeg state with maxMistakes 8",
			mistakes:    7,
			maxMistakes: 8,
			expected:    LeftLeg,
		},
		{
			name:        "RightLeg state with maxMistakes 8",
			mistakes:    8,
			maxMistakes: 8,
			expected:    RightLeg,
		},
	}

	assertInstance := assert.New(t)

	for _, tt := range tests {
		word := &Word{Word: "test", Hint: "test hint"}
		game := NewGame(word, tt.maxMistakes)
		game.mistakes = tt.mistakes

		assertInstance.Equal(tt.expected, game.State(), tt.name)
	}
}

func TestNewGame(t *testing.T) {
	tests := []struct {
		name        string
		word        *Word
		maxMistakes int
		expected    *Game
	}{
		{
			name:        "new game with single word",
			word:        &Word{Word: "apple", Hint: "A fruit"},
			maxMistakes: 5,
			expected: &Game{
				attempts:       0,
				mistakes:       0,
				maxMistakes:    5,
				word:           Word{Word: "apple", Hint: "A fruit"},
				correctLetters: map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
				used:           map[rune]bool{},
			},
		},
		{
			name:        "new game with word containing spaces",
			word:        &Word{Word: "hello world", Hint: "A greeting"},
			maxMistakes: 7,
			expected: &Game{
				attempts:       0,
				mistakes:       0,
				maxMistakes:    7,
				word:           Word{Word: "hello world", Hint: "A greeting"},
				correctLetters: map[rune]bool{'h': true, 'e': true, 'l': true, 'o': true, 'w': true, 'r': true, 'd': true, ' ': true},
				used:           map[rune]bool{' ': true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewGame(tt.word, tt.maxMistakes)
			assert.Equal(t, tt.expected.Attempts(), game.Attempts())
			assert.Equal(t, tt.expected.Mistakes(), game.Mistakes())
			assert.Equal(t, tt.expected.MaxMistakes(), game.MaxMistakes())
			assert.Equal(t, tt.expected.Used(), game.Used())
		})
	}
}

func TestPattern(t *testing.T) {
	tests := []struct {
		name           string
		word           *Word
		correctLetters map[rune]bool
		usedLetters    map[rune]bool
		expected       string
	}{
		{
			name:           "All letters guessed",
			word:           &Word{Word: "apple", Hint: "A fruit"},
			correctLetters: map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
			usedLetters:    map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
			expected:       "apple",
		},
		{
			name:           "No letters guessed",
			word:           &Word{Word: "apple", Hint: "A fruit"},
			correctLetters: map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
			usedLetters:    map[rune]bool{},
			expected:       "_____",
		},
		{
			name:           "Some letters guessed",
			word:           &Word{Word: "apple", Hint: "A fruit"},
			correctLetters: map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
			usedLetters:    map[rune]bool{'a': true, 'p': true},
			expected:       "app__",
		},
		{
			name:           "Word with spaces",
			word:           &Word{Word: "hello world", Hint: "A greeting"},
			correctLetters: map[rune]bool{'h': true, 'e': true, 'l': true, 'o': true, 'w': true, 'r': true, 'd': true, ' ': true},
			usedLetters:    map[rune]bool{'h': true, 'e': true, 'l': true, 'o': true, ' ': true},
			expected:       "hello _o_l_",
		},
		{
			name:           "Word with repeated letters",
			word:           &Word{Word: "banana", Hint: "A fruit"},
			correctLetters: map[rune]bool{'b': true, 'a': true, 'n': true},
			usedLetters:    map[rune]bool{'b': true, 'a': true},
			expected:       "ba_a_a",
		},
		{
			name: "Word with spaces",
			word: &Word{Word: "mother in law", Hint: "Family member"},
			correctLetters: map[rune]bool{
				'm': true, 'o': true, 't': true, 'h': true, 'e': true, 'r': true,
				'i': true, 'n': true, 'l': true, 'a': true, 'w': true, ' ': true,
			},
			usedLetters: map[rune]bool{'m': true, 'o': true, 't': true, 'h': true, 'e': true, 'r': true, 'i': true, 'n': true, ' ': true},
			expected:    "mother in ___",
		},
		{
			name:           "Word with all incorrect guesses",
			word:           &Word{Word: "kiwi", Hint: "A fruit"},
			correctLetters: map[rune]bool{'k': true, 'i': true, 'w': true},
			usedLetters:    map[rune]bool{'x': true, 'y': true, 'z': true},
			expected:       "____",
		},
	}

	assertInstance := assert.New(t)

	for _, tt := range tests {
		game := &Game{
			word:           *tt.word,
			correctLetters: tt.correctLetters,
			used:           tt.usedLetters,
		}
		game.used = tt.usedLetters
		assertInstance.Equal(tt.expected, game.Pattern(), tt.name)
	}
}

func TestGuess(t *testing.T) {
	tests := []struct {
		name             string
		attempts         int
		mistakes         int
		maxMistakes      int
		word             *Word
		correctLetters   map[rune]bool
		usedLetters      map[rune]bool
		guess            rune
		expectedAttempts int
		expectedMistakes int
		expectedUsed     map[rune]bool
	}{
		{
			name:             "Correct guess",
			attempts:         0,
			mistakes:         0,
			maxMistakes:      5,
			word:             &Word{Word: "apple", Hint: "A fruit"},
			correctLetters:   map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
			usedLetters:      map[rune]bool{},
			guess:            'a',
			expectedAttempts: 1,
			expectedMistakes: 0,
			expectedUsed:     map[rune]bool{'a': true},
		},
		{
			name:             "Incorrect guess",
			attempts:         0,
			mistakes:         0,
			maxMistakes:      5,
			word:             &Word{Word: "apple", Hint: "A fruit"},
			correctLetters:   map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
			usedLetters:      map[rune]bool{},
			guess:            'b',
			expectedAttempts: 1,
			expectedMistakes: 1,
			expectedUsed:     map[rune]bool{'b': true},
		},
		{
			name:             "Repeated correct guess",
			attempts:         1,
			mistakes:         0,
			maxMistakes:      5,
			word:             &Word{Word: "apple", Hint: "A fruit"},
			correctLetters:   map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
			usedLetters:      map[rune]bool{'a': true},
			guess:            'a',
			expectedAttempts: 1,
			expectedMistakes: 0,
			expectedUsed:     map[rune]bool{'a': true},
		},
		{
			name:             "Repeated incorrect guess",
			attempts:         1,
			mistakes:         1,
			maxMistakes:      5,
			word:             &Word{Word: "apple", Hint: "A fruit"},
			correctLetters:   map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
			usedLetters:      map[rune]bool{'b': true},
			guess:            'b',
			expectedAttempts: 1,
			expectedMistakes: 1,
			expectedUsed:     map[rune]bool{'b': true},
		},
		{
			name:             "Correct guess in multi-word",
			attempts:         0,
			mistakes:         0,
			maxMistakes:      7,
			word:             &Word{Word: "hello world", Hint: "A greeting"},
			correctLetters:   map[rune]bool{'h': true, 'e': true, 'l': true, 'o': true, 'w': true, 'r': true, 'd': true, ' ': true},
			usedLetters:      map[rune]bool{' ': true},
			guess:            'h',
			expectedAttempts: 1,
			expectedMistakes: 0,
			expectedUsed:     map[rune]bool{'h': true, ' ': true},
		},
		{
			name:             "Incorrect guess in multi-word",
			attempts:         0,
			mistakes:         0,
			maxMistakes:      7,
			word:             &Word{Word: "hello world", Hint: "A greeting"},
			correctLetters:   map[rune]bool{'h': true, 'e': true, 'l': true, 'o': true, 'w': true, 'r': true, 'd': true, ' ': true},
			usedLetters:      map[rune]bool{' ': true},
			guess:            'x',
			expectedAttempts: 1,
			expectedMistakes: 1,
			expectedUsed:     map[rune]bool{'x': true, ' ': true},
		},
		{
			name:             "Repeated correct guess in multi-word",
			attempts:         1,
			mistakes:         0,
			maxMistakes:      7,
			word:             &Word{Word: "hello world", Hint: "A greeting"},
			correctLetters:   map[rune]bool{'h': true, 'e': true, 'l': true, 'o': true, 'w': true, 'r': true, 'd': true, ' ': true},
			usedLetters:      map[rune]bool{'h': true, ' ': true},
			guess:            'h',
			expectedAttempts: 1,
			expectedMistakes: 0,
			expectedUsed:     map[rune]bool{'h': true, ' ': true},
		},
		{
			name:             "Repeated incorrect guess in multi-word",
			attempts:         1,
			mistakes:         1,
			maxMistakes:      7,
			word:             &Word{Word: "hello world", Hint: "A greeting"},
			correctLetters:   map[rune]bool{'h': true, 'e': true, 'l': true, 'o': true, 'w': true, 'r': true, 'd': true, ' ': true},
			usedLetters:      map[rune]bool{'x': true, ' ': true},
			guess:            'x',
			expectedAttempts: 1,
			expectedMistakes: 1,
			expectedUsed:     map[rune]bool{'x': true, ' ': true},
		},
	}

	assertInstance := assert.New(t)

	for _, tt := range tests {
		game := &Game{
			attempts:       tt.attempts,
			mistakes:       tt.mistakes,
			maxMistakes:    tt.maxMistakes,
			word:           *tt.word,
			correctLetters: tt.correctLetters,
			used:           tt.usedLetters,
		}
		game.Guess(tt.guess)
		assertInstance.Equal(tt.expectedAttempts, game.Attempts(), tt.name)
		assertInstance.Equal(tt.expectedMistakes, game.Mistakes(), tt.name)
		assertInstance.Equal(tt.expectedUsed, game.Used(), tt.name)
	}
}

func TestIsWin(t *testing.T) {
	tests := []struct {
		name           string
		attempts       int
		mistakes       int
		maxMistakes    int
		word           *Word
		correctLetters map[rune]bool
		usedLetters    map[rune]bool
		expected       bool
	}{
		{
			name:           "Win with all letters guessed",
			attempts:       5,
			mistakes:       1,
			maxMistakes:    6,
			word:           &Word{Word: "apple", Hint: "A fruit"},
			correctLetters: map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
			usedLetters:    map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
			expected:       true,
		},
		{
			name:           "Not win with some letters guessed",
			attempts:       5,
			mistakes:       1,
			maxMistakes:    6,
			word:           &Word{Word: "apple", Hint: "A fruit"},
			correctLetters: map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
			usedLetters:    map[rune]bool{'a': true, 'p': true},
			expected:       false,
		},
		{
			name:           "Not win with no letters guessed",
			attempts:       5,
			mistakes:       1,
			maxMistakes:    6,
			word:           &Word{Word: "apple", Hint: "A fruit"},
			correctLetters: map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
			usedLetters:    map[rune]bool{},
			expected:       false,
		},
		{
			name:           "Lose with max mistakes",
			attempts:       6,
			mistakes:       6,
			maxMistakes:    6,
			word:           &Word{Word: "apple", Hint: "A fruit"},
			correctLetters: map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
			usedLetters:    map[rune]bool{'z': true, 'y': true, 'x': true, 'w': true, 'v': true, 'u': true},
			expected:       false,
		},
		{
			name:           "Win with multi-word all letters guessed",
			attempts:       10,
			mistakes:       2,
			maxMistakes:    7,
			word:           &Word{Word: "hello world", Hint: "A greeting"},
			correctLetters: map[rune]bool{'h': true, 'e': true, 'l': true, 'o': true, 'w': true, 'r': true, 'd': true, ' ': true},
			usedLetters:    map[rune]bool{'h': true, 'e': true, 'l': true, 'o': true, 'w': true, 'r': true, 'd': true, ' ': true},
			expected:       true,
		},
		{
			name:           "Not win with multi-word some letters guessed",
			attempts:       10,
			mistakes:       2,
			maxMistakes:    7,
			word:           &Word{Word: "hello world", Hint: "A greeting"},
			correctLetters: map[rune]bool{'h': true, 'e': true, 'l': true, 'o': true, 'w': true, 'r': true, 'd': true, ' ': true},
			usedLetters:    map[rune]bool{'h': true, 'e': true, 'l': true, 'o': true, ' ': true},
			expected:       false,
		},
		{
			name:           "Not win with multi-word no letters guessed",
			attempts:       10,
			mistakes:       5,
			maxMistakes:    7,
			word:           &Word{Word: "hello world", Hint: "A greeting"},
			correctLetters: map[rune]bool{'h': true, 'e': true, 'l': true, 'o': true, 'w': true, 'r': true, 'd': true, ' ': true},
			usedLetters: map[rune]bool{
				' ': true, 'h': true, 'e': true, 'l': true, 'o': true, 'w': true,
				'y': true, 'z': true, 'v': true, 'u': true, 't': true,
			},
			expected: false,
		},
	}

	assertInstance := assert.New(t)

	for _, tt := range tests {
		game := &Game{
			attempts:       tt.attempts,
			mistakes:       tt.mistakes,
			maxMistakes:    tt.maxMistakes,
			word:           *tt.word,
			correctLetters: tt.correctLetters,
			used:           tt.usedLetters,
		}
		assertInstance.Equal(tt.expected, game.IsWin(), tt.name)
	}
}

func TestIsLose(t *testing.T) {
	tests := []struct {
		name           string
		attempts       int
		mistakes       int
		maxMistakes    int
		word           *Word
		correctLetters map[rune]bool
		usedLetters    map[rune]bool
		expected       bool
	}{
		{
			name:           "Lose with max mistakes",
			attempts:       6,
			mistakes:       6,
			maxMistakes:    6,
			word:           &Word{Word: "apple", Hint: "A fruit"},
			correctLetters: map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
			usedLetters:    map[rune]bool{'z': true, 'y': true, 'x': true, 'w': true, 'v': true, 'u': true},
			expected:       true,
		},
		{
			name:           "Not lose with some mistakes",
			attempts:       5,
			mistakes:       3,
			maxMistakes:    6,
			word:           &Word{Word: "apple", Hint: "A fruit"},
			correctLetters: map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
			usedLetters:    map[rune]bool{'a': true, 'p': true},
			expected:       false,
		},
		{
			name:           "Not lose with no mistakes",
			attempts:       0,
			mistakes:       0,
			maxMistakes:    6,
			word:           &Word{Word: "apple", Hint: "A fruit"},
			correctLetters: map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
			usedLetters:    map[rune]bool{},
			expected:       false,
		},
		{
			name:           "Lose with max mistakes in multi-word",
			attempts:       10,
			mistakes:       7,
			maxMistakes:    7,
			word:           &Word{Word: "hello world", Hint: "A greeting"},
			correctLetters: map[rune]bool{'h': true, 'e': true, 'l': true, 'o': true, 'w': true, 'r': true, 'd': true, ' ': true},
			usedLetters: map[rune]bool{
				' ': true, 'h': true, 'e': true, 'l': true, 'o': true, 'w': true,
				'y': true, 'z': true, 'v': true, 'u': true, 't': true, 'k': true, 'j': true,
			},
			expected: true,
		},
	}

	assertInstance := assert.New(t)

	for _, tt := range tests {
		game := &Game{
			attempts:       tt.attempts,
			mistakes:       tt.mistakes,
			maxMistakes:    tt.maxMistakes,
			word:           *tt.word,
			correctLetters: tt.correctLetters,
			used:           tt.usedLetters,
		}
		assertInstance.Equal(tt.expected, game.IsLose(), tt.name)
	}
}

func TestIsFinished(t *testing.T) {
	tests := []struct {
		name           string
		attempts       int
		mistakes       int
		maxMistakes    int
		word           *Word
		correctLetters map[rune]bool
		usedLetters    map[rune]bool
		expected       bool
	}{
		{
			name:           "Game finished with win",
			attempts:       5,
			mistakes:       1,
			maxMistakes:    6,
			word:           &Word{Word: "apple", Hint: "A fruit"},
			correctLetters: map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
			usedLetters:    map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
			expected:       true,
		},
		{
			name:           "Game finished with lose",
			attempts:       6,
			mistakes:       6,
			maxMistakes:    6,
			word:           &Word{Word: "apple", Hint: "A fruit"},
			correctLetters: map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
			usedLetters:    map[rune]bool{'z': true, 'y': true, 'x': true, 'w': true, 'v': true, 'u': true},
			expected:       true,
		},
		{
			name:           "Game not finished with some mistakes",
			attempts:       5,
			mistakes:       3,
			maxMistakes:    6,
			word:           &Word{Word: "apple", Hint: "A fruit"},
			correctLetters: map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
			usedLetters:    map[rune]bool{'a': true, 'p': true},
			expected:       false,
		},
		{
			name:           "Game not finished with no mistakes",
			attempts:       0,
			mistakes:       0,
			maxMistakes:    6,
			word:           &Word{Word: "apple", Hint: "A fruit"},
			correctLetters: map[rune]bool{'a': true, 'p': true, 'l': true, 'e': true},
			usedLetters:    map[rune]bool{},
			expected:       false,
		},
		{
			name:           "Game finished with win in multi-word",
			attempts:       10,
			mistakes:       2,
			maxMistakes:    7,
			word:           &Word{Word: "hello world", Hint: "A greeting"},
			correctLetters: map[rune]bool{'h': true, 'e': true, 'l': true, 'o': true, 'w': true, 'r': true, 'd': true, ' ': true},
			usedLetters: map[rune]bool{
				'h': true, 'e': true, 'l': true, 'o': true, 'w': true, 'r': true,
				'd': true, 'q': true, 's': true, ' ': true,
			},
			expected: true,
		},
	}

	assertInstance := assert.New(t)

	for _, tt := range tests {
		game := &Game{
			attempts:       tt.attempts,
			mistakes:       tt.mistakes,
			maxMistakes:    tt.maxMistakes,
			word:           *tt.word,
			correctLetters: tt.correctLetters,
			used:           tt.usedLetters,
		}
		assertInstance.Equal(tt.expected, game.IsFinished(), tt.name)
	}
}
