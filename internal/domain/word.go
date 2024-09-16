package domain

import "log/slog"

type Word struct {
	Word string `json:"word"`
	Hint string `json:"hint"`
}

func (w *Word) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("word", w.Word),
	)
}
