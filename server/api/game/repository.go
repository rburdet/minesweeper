package game

type Repository interface {
	SaveGame(name string, game Game)
	GetGame(name string) Game
	DeleteGame(name string)
	Exists(name string) bool
}
