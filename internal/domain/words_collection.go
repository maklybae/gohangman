package domain

type WordsCollection struct {
	Creator     string     `json:"creator"`
	Description string     `json:"description"`
	Categories  []Category `json:"categories"`
}
