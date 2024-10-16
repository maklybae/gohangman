package domain

type GameInputer interface {
	GetLetter() (letter rune, err error)
}

type InputerError struct {
	Message    string
	InnerError error
}

func (e *InputerError) Error() string {
	return e.Message
}

func (e *InputerError) Unwrap() error {
	return e.InnerError
}
