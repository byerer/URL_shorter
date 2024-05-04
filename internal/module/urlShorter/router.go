package urlShorter

import (
	"URL_shorter/internal/global/errs"
	"github.com/gin-gonic/gin"
)

func (u *URLShorter) InitRouter(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/shorten", Shorten)
	r.GET("/:shorten", Redirect)
}

func Shorten(c *gin.Context) {
	req := struct {
		URL string `json:"url" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		errs.Fail(c, errs.InvalidRequest.WithOrigin(err))
		return
	}
	err := CreateUrl(req.URL)
	if err != nil {
		errs.Fail(c, errs.DatabaseError.WithOrigin(err))
		return
	}
	errs.Success(c, "create success")
}

func Redirect(c *gin.Context) {
	shorten := c.Param("shorten")
	url, err := GetUrl(shorten)
	if err != nil {
		errs.Fail(c, errs.DatabaseError.WithOrigin(err))
		return
	}
	c.Redirect(301, url)
}
