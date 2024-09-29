package domain

import "log/slog"

type WordJSON struct {
	Word string `json:"word"`
	Hint string `json:"hint"`
}

func (w *WordJSON) ToDomain() *Word {
	return &Word{
		Word: w.Word,
		Hint: w.Hint,
	}
}

type Word struct {
	Word string
	Hint string
}

func (w *Word) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("word", w.Word),
	)
}
