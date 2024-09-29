package domain

import (
	"fmt"
	"log/slog"
)

type CategoryJSON struct {
	Name        string     `json:"name"`
	EasyWords   []WordJSON `json:"easy"`
	MediumWords []WordJSON `json:"medium"`
	HardWords   []WordJSON `json:"hard"`
}

func (c *CategoryJSON) ToDomain() *Category {
	easyWords := make([]Word, 0, len(c.EasyWords))
	mediumWords := make([]Word, 0, len(c.MediumWords))
	hardWords := make([]Word, 0, len(c.HardWords))

	for _, word := range c.EasyWords {
		easyWords = append(easyWords, *word.ToDomain())
	}

	for _, word := range c.MediumWords {
		mediumWords = append(mediumWords, *word.ToDomain())
	}

	for _, word := range c.HardWords {
		hardWords = append(hardWords, *word.ToDomain())
	}

	return &Category{
		Name:        c.Name,
		EasyWords:   easyWords,
		MediumWords: mediumWords,
		HardWords:   hardWords,
	}
}

type Category struct {
	Name        string
	EasyWords   []Word
	MediumWords []Word
	HardWords   []Word
}

func (c *Category) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("name", c.Name),
		slog.Int("easy words count", len(c.EasyWords)),
		slog.Int("medium words count", len(c.MediumWords)),
		slog.Int("hard words count", len(c.HardWords)),
	)
}

type BadCategoryError struct {
	Message string
}

func (e *BadCategoryError) Error() string {
	return fmt.Sprintf("bad category: %s", e.Message)
}
