package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"minesweeper/api/game"
	"minesweeper/api/health"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api")
	loadHealth(api)
	loadGame(api)
	router.Run()
}

func loadHealth(router *gin.RouterGroup) {
	handler := &health.HealthHandler{}
	router.GET("/health", handler.Health)
}

func loadGame(router *gin.RouterGroup) {
	db := game.NewMemDbImpl()
	service := game.NewGameService(db)
	handler := game.NewGameHandler(service)
	router.POST("/game", handler.CreateGame)
	router.GET("/game/:name", handler.GetGame)
	router.POST("/game/:name/click", handler.Click)
}
