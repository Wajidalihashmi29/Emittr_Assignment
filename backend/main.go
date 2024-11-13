package main

import (
	"exploding_kitten/controllers"
	"exploding_kitten/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitRedis()
	router := gin.Default()

	router.POST("/start-game", controllers.StartGame)
	router.POST("/draw-card", controllers.DrawCard)
	router.GET("/leaderboard", controllers.GetLeaderboard)

	router.Run(":8080")
}
