package server

import (
	"URL_shorter/internal/global/database"
	"github.com/gin-gonic/gin"
)

func Init() {
	database.Init()
	database.Ping()
}

func Run() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
