package server

import (
	"URL_shorter/internal/global/database/pgsql"
	"URL_shorter/internal/global/database/redis"
	"URL_shorter/internal/module"
	"github.com/gin-gonic/gin"
)

func Init() {
	pgsql.Init()
	redis.Init()
	redis.Ping()
}

func Run() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	apiGroup := r.Group("/api")
	for _, m := range module.Modules {
		m.InitRouter(apiGroup)
	}
	r.Run()
}

func Close() {
	redis.Close()
}
