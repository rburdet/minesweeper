package game

import (
	"github.com/Pallinder/go-randomdata"
	"math/rand"
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
	Rows    int        `json:"rows" binding:"gt=2,lt=100"`
	Columns int        ` json:"columns" binding:"gt=2,lt=100"`
	Mines   int        `json:"mines" binding:"gt=1,lt=100"`
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

func (g *Game) fillCell() {
	boardSize := g.Rows * g.Columns
	target := rand.Intn(boardSize)
	row := target / g.Rows
	column := target % g.Columns
	cell := &g.Board[row][column]
	if cell.HasMine {
		g.fillCell()
	} else {
		cell.HasMine = true
	}
}

func (g *Game) sumAdjacents(i int, j int) {
	for z := i - 1; z < i+2; z++ {
		if z < 0 || z > g.Rows-1 {
			continue
		}
		for w := j - 1; w < j+2; w++ {
			if w < 0 || w > g.Columns-1 {
				continue
			}
			if z == i && w == j {
				continue
			}
			g.Board[z][w].AdjacentMines++
		}
	}
}

func (g *Game) fillBoard() {
	g.Board = make([][]Cell, g.Rows)
	for i := 0; i < g.Rows; i++ {
		g.Board[i] = make([]Cell, g.Columns)
	}

	for i := 0; i < g.Mines; i++ {
		g.fillCell()
	}

	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Columns; j++ {
			if g.Board[i][j].HasMine {
				g.sumAdjacents(i, j)
			}
		}
	}
}

func (g *Game) floodFill(i int, j int) {
	if i >= 0 && j >= 0 && i < g.Rows && j < g.Columns {
		cell := &g.Board[i][j]
		if cell.IsClicked {
			return
		}
		cell.IsClicked = true
		if cell.AdjacentMines == 0 {
			g.floodFill(i-1, j)
			g.floodFill(i-1, j-1)
			g.floodFill(i+1, j)
			g.floodFill(i+1, j+1)
			g.floodFill(i, j-1)
			g.floodFill(i+1, j-1)
			g.floodFill(i, j+1)
			g.floodFill(i-1, j+1)
		} else {
			return
		}
	}
}

func (g *Game) checkGameEnded() {
	var notOpenedCells int
	var discoveredMines int
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Columns; j++ {
			cell := g.Board[i][j]
			if !cell.IsClicked {
				notOpenedCells++
			}
			if cell.Status == Flag && cell.HasMine {
				discoveredMines++
			}
		}
	}
	if notOpenedCells == g.Mines || discoveredMines == g.Mines {
		g.Status = Won
	}
}

func (g *Game) Click(i int, j int, action string) {
	cell := &g.Board[i][j]
	if action == "clean" {
		cell.Status = ""
	} else if action == "flag" {
		cell.Status = Flag
	} else if action == "question" {
		cell.Status = QuestionMark
	} else {
		if cell.HasMine {
			g.Status = Lost
			return
		}
		g.floodFill(i, j)
	}

	g.checkGameEnded()
}
