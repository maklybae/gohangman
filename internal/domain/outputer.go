package domain

type GameOutputer interface {
	ShowPattern(pattern string)
	ShowState(state State)
}
