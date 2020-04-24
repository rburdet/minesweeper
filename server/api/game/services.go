package game

type GameService struct {
	repository Repository
}

func NewGameService(repository Repository) *GameService {
	return &GameService{repository}
}

func (gs *GameService) SaveGame(g Game) {
	gs.repository.SaveGame(g.Name, g)
}

func (gs *GameService) GetGame(name string) Game {
	return gs.repository.GetGame(name)
}

func (gs *GameService) Exists(name string) bool {
	return gs.repository.Exists(name)
}
