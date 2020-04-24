package game

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ClickAction struct {
	I      int
	J      int
	Action string `binding:oneof="flag question"`
}

type GameHandler struct {
	service *GameService
}

func NewGameHandler(s *GameService) *GameHandler {
	return &GameHandler{s}
}

func (gh *GameHandler) CreateGame(c *gin.Context) {
	newGame := NewGame()
	if err := c.ShouldBindJSON(&newGame); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newGame.Rows*newGame.Columns < newGame.Mines {
		c.JSON(http.StatusForbidden, gin.H{"error": "too many mines for that board size"})
		return
	}

	newGame.fillBoard()
	if gh.service.Exists(newGame.Name) {
		c.JSON(http.StatusForbidden, gin.H{"error": "already exists"})
		return
	}
	gh.service.SaveGame(newGame)

	c.JSON(http.StatusOK, gin.H{"id": newGame.Name})
}

func (gh *GameHandler) GetGame(c *gin.Context) {
	name := c.Param("name")
	savedGame := gh.service.GetGame(name)

	c.JSON(http.StatusOK, savedGame)
}
