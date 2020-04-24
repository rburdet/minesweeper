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

	if savedGame.Name == "" {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, savedGame)
}

func (gh *GameHandler) Click(c *gin.Context) {
	var click ClickAction
	if err := c.ShouldBindJSON(&click); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	name := c.Param("name")

	savedGame := gh.service.GetGame(name)
	if savedGame.Status != Playing {
		c.JSON(http.StatusForbidden, gin.H{"error": "game is over"})

		return
	}
	savedGame.Click(click.I, click.J, click.Action)
	gh.service.SaveGame(savedGame)

	c.JSON(http.StatusOK, savedGame)
}
