package climenu

type ExitError struct {
}

func (e *ExitError) Error() string {
	return "menu exit command"
}
