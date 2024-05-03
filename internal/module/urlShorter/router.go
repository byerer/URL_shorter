package urlShorter

import "github.com/gin-gonic/gin"

func (u *URLShorter) InitRouter(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/shorten", Shorten)
}

func Shorten(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "shorten",
	})
}
