package urlShorter

import (
	"URL_shorter/internal/global/errs"
	"URL_shorter/internal/model"
	"URL_shorter/service/shorter"
	"URL_shorter/tools"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (u *URLShorter) InitRouter(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/shorter", Shorten)
	r.GET("/:shorter", Redirect)
}

func Shorten(c *gin.Context) {
	req := struct {
		URL string `json:"url" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		errs.Fail(c, errs.InvalidRequest.WithOrigin(err))
		return
	}
	url := &model.Url{
		LongURL:  req.URL,
		ShortURL: shorter.Shorten(req.URL),
	}
	err := CreateUrl(url)
	for tools.IsDuplicateKeyError(err) {
		fmt.Println("Duplicate key error, retrying")
		url.ShortURL += shorter.Add1(1)
		err = CreateUrl(url)
	}
	if err != nil {
		errs.Fail(c, errs.DatabaseError.WithOrigin(err))
		return
	}
	errs.Success(c, c.Request.Host+"/"+url.ShortURL)
}

func Redirect(c *gin.Context) {
	shorten := c.Param("shorter")
	url, err := GetUrl(shorten)
	if err != nil {
		errs.Fail(c, errs.DatabaseError.WithOrigin(err))
		return
	}
	c.Redirect(302, url)
}
