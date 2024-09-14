package domain

import "fmt"

type WordsCollection struct {
	Creator     string     `json:"creator"`
	Description string     `json:"description"`
	Categories  []Category `json:"categories"`
}

type BadWordsCollectionError struct {
	Message string
}

func (e *BadWordsCollectionError) Error() string {
	return fmt.Sprintf("bad words collection: %s", e.Message)
}
