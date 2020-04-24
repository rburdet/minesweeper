package game

import (
	"github.com/Pallinder/go-randomdata"
	"strings"
)

type GameStatus string
type CellStatus string

const (
	Flag         CellStatus = "flag"
	QuestionMark            = "question"
)

const (
	Won     GameStatus = "won"
	Lost               = "lost"
	Playing            = "playing"
)

type Cell struct {
	HasMine       bool       `json:"hasMine"`
	Status        CellStatus `json:"status"`
	IsClicked     bool       `json:"clicked"`
	AdjacentMines int        `json:"adjancetMines"`
}

type Game struct {
	Rows    int        `json:"rows" binding:"gt=2"`
	Columns int        ` json:"columns" binding:"gt=2"`
	Mines   int        `json:"mines" binding:"gt=1"`
	Name    string     `json:"name"`
	Status  GameStatus `json:"status"`
	Board   [][]Cell   `json:"board"`
}

func NewGame() Game {
	newGame := Game{}
	newGame.Name = strings.ToLower(randomdata.SillyName())
	newGame.Status = Playing

	return newGame
}
