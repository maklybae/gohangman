package domain

import (
	"fmt"
	"log/slog"
)

type WordsCollection struct {
	Creator     string     `json:"creator"`
	Description string     `json:"description"`
	Categories  []Category `json:"categories"`
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
