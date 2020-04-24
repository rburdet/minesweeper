package game

import (
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
