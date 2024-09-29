package application_test

import (
	"io"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"makly/hangman/internal/application"
	applicationMocks "makly/hangman/internal/application/mocks"
	"makly/hangman/internal/domain"
	domainMocks "makly/hangman/internal/domain/mocks"
)

func TestChoiceDifficulty(t *testing.T) {
	log.SetOutput(io.Discard)

	const iterations = 1000

	difficultyCount := make(map[domain.Difficulty]int)

	for i := 0; i < iterations; i++ {
		difficulty, err := application.ChoiceDifficulty()

		assert.NoError(t, err)

		difficultyCount[difficulty]++
	}

	for i := 0; i < int(domain.DifficultyCount); i++ {
		assert.Greater(t, difficultyCount[domain.Difficulty(i)], 0, "Difficulty %d was never chosen", i)
	}
}

func TestChoiceCategory(t *testing.T) {
	log.SetOutput(io.Discard)

	const iterations = 1000

	categories := []domain.Category{
		{Name: "Category1"},
		{Name: "Category2"},
		{Name: "Category3"},
		{Name: "Category4"},
		{Name: "Category5"},
		{Name: "Category6"},
		{Name: "Category7"},
		{Name: "Category8"},
		{Name: "Category9"},
		{Name: "Category10"},
	}

	categoryCount := make(map[string]int)

	for i := 0; i < iterations; i++ {
		category, err := application.ChoiceCategory(categories)

		assert.NoError(t, err)

		categoryCount[category.Name]++
	}

	for _, category := range categories {
		assert.Greater(t, categoryCount[category.Name], 0, "Category %s was never chosen", category.Name)
	}
}

func TestChoiceWord(t *testing.T) {
	log.SetOutput(io.Discard)

	rd := &application.RandomDefault{}

	const iterations = 1000

	category := &domain.Category{
		Name: "Category",
		EasyWords: []domain.Word{
			{Word: "easy1", Hint: "hint_easy1"},
			{Word: "easy2", Hint: "hint_easy2"},
			{Word: "easy3", Hint: "hint_easy3"},
			{Word: "easy4", Hint: "hint_easy4"},
			{Word: "easy5", Hint: "hint_easy5"},
			{Word: "easy6", Hint: "hint_easy6"},
			{Word: "easy7", Hint: "hint_easy7"},
			{Word: "easy8", Hint: "hint_easy8"},
			{Word: "easy9", Hint: "hint_easy9"},
			{Word: "easy10", Hint: "hint_easy10"},
			{Word: "easy11", Hint: "hint_easy11"},
			{Word: "easy12", Hint: "hint_easy12"},
			{Word: "easy13", Hint: "hint_easy13"},
			{Word: "easy14", Hint: "hint_easy14"},
			{Word: "easy15", Hint: "hint_easy15"},
			{Word: "easy16", Hint: "hint_easy16"},
			{Word: "easy17", Hint: "hint_easy17"},
			{Word: "easy18", Hint: "hint_easy18"},
			{Word: "easy19", Hint: "hint_easy19"},
			{Word: "easy20", Hint: "hint_easy20"},
		},
		MediumWords: []domain.Word{
			{Word: "medium1", Hint: "hint_medium1"},
			{Word: "medium2", Hint: "hint_medium2"},
			{Word: "medium3", Hint: "hint_medium3"},
			{Word: "medium4", Hint: "hint_medium4"},
			{Word: "medium5", Hint: "hint_medium5"},
			{Word: "medium6", Hint: "hint_medium6"},
			{Word: "medium7", Hint: "hint_medium7"},
			{Word: "medium8", Hint: "hint_medium8"},
			{Word: "medium9", Hint: "hint_medium9"},
			{Word: "medium10", Hint: "hint_medium10"},
			{Word: "medium11", Hint: "hint_medium11"},
			{Word: "medium12", Hint: "hint_medium12"},
			{Word: "medium13", Hint: "hint_medium13"},
			{Word: "medium14", Hint: "hint_medium14"},
			{Word: "medium15", Hint: "hint_medium15"},
			{Word: "medium16", Hint: "hint_medium16"},
			{Word: "medium17", Hint: "hint_medium17"},
			{Word: "medium18", Hint: "hint_medium18"},
			{Word: "medium19", Hint: "hint_medium19"},
			{Word: "medium20", Hint: "hint_medium20"},
		},
		HardWords: []domain.Word{
			{Word: "hard1", Hint: "hint_hard1"},
			{Word: "hard2", Hint: "hint_hard2"},
			{Word: "hard3", Hint: "hint_hard3"},
			{Word: "hard4", Hint: "hint_hard4"},
			{Word: "hard5", Hint: "hint_hard5"},
			{Word: "hard6", Hint: "hint_hard6"},
			{Word: "hard7", Hint: "hint_hard7"},
			{Word: "hard8", Hint: "hint_hard8"},
			{Word: "hard9", Hint: "hint_hard9"},
			{Word: "hard10", Hint: "hint_hard10"},
			{Word: "hard11", Hint: "hint_hard11"},
			{Word: "hard12", Hint: "hint_hard12"},
			{Word: "hard13", Hint: "hint_hard13"},
			{Word: "hard14", Hint: "hint_hard14"},
			{Word: "hard15", Hint: "hint_hard15"},
			{Word: "hard16", Hint: "hint_hard16"},
			{Word: "hard17", Hint: "hint_hard17"},
			{Word: "hard18", Hint: "hint_hard18"},
			{Word: "hard19", Hint: "hint_hard19"},
			{Word: "hard20", Hint: "hint_hard20"},
		},
	}

	wordCount := make(map[string]int)

	for i := 0; i < iterations; i++ {
		word, err := rd.ChoiceWord(category, domain.EasyDifficulty)

		assert.NoError(t, err)

		wordCount[word.Word]++
	}

	for _, word := range category.EasyWords {
		assert.Greater(t, wordCount[word.Word], 0, "Word %s was never chosen", word.Word)
	}

	wordCount = make(map[string]int)

	for i := 0; i < iterations; i++ {
		word, err := rd.ChoiceWord(category, domain.MediumDifficulty)

		assert.NoError(t, err)

		wordCount[word.Word]++
	}

	for _, word := range category.MediumWords {
		assert.Greater(t, wordCount[word.Word], 0, "Word %s was never chosen", word.Word)
	}

	wordCount = make(map[string]int)

	for i := 0; i < iterations; i++ {
		word, err := rd.ChoiceWord(category, domain.HardDifficulty)

		assert.NoError(t, err)

		wordCount[word.Word]++
	}

	for _, word := range category.HardWords {
		assert.Greater(t, wordCount[word.Word], 0, "Word %s was never chosen", word.Word)
	}
}

func TestRunGameSession(t *testing.T) {
	log.SetOutput(io.Discard)

	mockWordRandomizer := &applicationMocks.WordRandomizer{}
	mockInputer := &domainMocks.GameInputer{}
	mockOutputer := &domainMocks.GameOutputer{}

	mockWordRandomizer.On("ChoiceWord", mock.Anything, mock.Anything).Return(&domain.Word{Word: "test word"}, nil)

	mockInputer.On("GetLetter").Return(rune('t'), nil).Once()
	mockInputer.On("GetLetter").Return(rune('e'), nil).Once()
	mockInputer.On("GetLetter").Return(rune('s'), nil).Once()
	mockInputer.On("GetLetter").Return(rune('w'), nil).Once()
	mockInputer.On("GetLetter").Return(rune('o'), nil).Once()
	mockInputer.On("GetLetter").Return(rune('r'), nil).Once()
	mockInputer.On("GetLetter").Return(rune('d'), nil).Once()

	// Check number of updates - 1 initial + 7 letters + 1 final
	mockOutputer.On("ShowGame", mock.Anything).Return().Times(1 + 7 + 1)
	mockOutputer.On("ShowGameResult", mock.Anything).Return().Once()

	err := application.RunGameSession(nil, domain.UnknownDifficulty, 6, mockInputer, mockOutputer, mockWordRandomizer)
	// Return from infinite game-loop check
	assert.Nil(t, err)
}
