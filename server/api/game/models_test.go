package game

import (
	"reflect"
	"testing"
)

func TestCreateNewGame(t *testing.T) {
	g := NewGame()
	if g.Status != Playing {
		t.Errorf("Status should be playing but is %s", g.Status)
	}

	if g.Name == "" {
		t.Errorf("Name should be not empty but is %s", g.Name)
	}
}

func TestFillBoard(t *testing.T) {
	t.Run("Mines", testMinesCreation)
	t.Run("Adjacents", testAdjacents)
}

func testMinesCreation(t *testing.T) {
	g := NewGame()
	g.Mines = 10
	g.Rows = 10
	g.Columns = 10

	minesCounter := 0
	g.fillBoard()

	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Columns; j++ {
			if g.Board[i][j].HasMine {
				minesCounter++
			}
		}
	}

	if minesCounter != g.Mines {
		t.Errorf("There should be %d mines, but there are %d", g.Mines, minesCounter)
	}
}

func testAdjacents(t *testing.T) {
	g := NewGame()
	g.Rows = 3
	g.Columns = 3
	g.Board = [][]Cell{
		{{true, "", false, 0}, {false, "", false, 0}, {false, "", false, 0}},
		{{false, "", false, 0}, {false, "", false, 0}, {false, "", false, 0}},
		{{false, "", false, 0}, {true, "", false, 0}, {false, "", false, 0}},
	}
	g.sumAdjacents(0, 0)
	g.sumAdjacents(2, 1)

	expected := [][]Cell{
		{{true, "", false, 0}, {false, "", false, 1}, {false, "", false, 0}},
		{{false, "", false, 2}, {false, "", false, 2}, {false, "", false, 1}},
		{{false, "", false, 1}, {true, "", false, 0}, {false, "", false, 1}},
	}

	if !reflect.DeepEqual(g.Board, expected) {
		t.Errorf("Error calculating adjacents \nexpected %v \nbut got %v", g.Board, expected)
	}

}
