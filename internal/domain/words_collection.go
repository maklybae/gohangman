package domain

import (
	"fmt"
	"log/slog"
)

type WordsCollectionJSON struct {
	Creator     string         `json:"creator"`
	Description string         `json:"description"`
	Categories  []CategoryJSON `json:"categories"`
}

func (w *WordsCollectionJSON) ToDomain() *WordsCollection {
	categories := make([]Category, 0, len(w.Categories))

	for _, category := range w.Categories {
		categories = append(categories, *category.ToDomain())
	}

	return &WordsCollection{
		Creator:     w.Creator,
		Description: w.Description,
		Categories:  categories,
	}
}

type WordsCollection struct {
	Creator     string
	Description string
	Categories  []Category
}

func (w *WordsCollection) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("creator", w.Creator),
		slog.String("description", w.Description),
		slog.Int("categories count", len(w.Categories)),
	)
}

type BadWordsCollectionError struct {
	Message string
}

func (e *BadWordsCollectionError) Error() string {
	return fmt.Sprintf("bad words collection: %s", e.Message)
}
