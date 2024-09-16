package domain

import (
	"fmt"
	"log/slog"
)

type Category struct {
	Name        string `json:"name"`
	EasyWords   []Word `json:"easy"`
	MediumWords []Word `json:"medium"`
	HardWords   []Word `json:"hard"`
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
