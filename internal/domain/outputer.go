package domain

type GameOutputer interface {
	ShowGame(game *Game)
	ShowGameResult(game *Game)
	ShowInputError(err error)
}
