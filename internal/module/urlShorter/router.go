package urlShorter

import (
	"URL_shorter/internal/global/errs"
	"URL_shorter/internal/model"
	"URL_shorter/internal/module/urlShorter/dao/pgsql"
	"URL_shorter/internal/module/urlShorter/dao/redis"
	"URL_shorter/service/shorter"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"time"
)

func (u *ModelURL) InitRouter(r *gin.RouterGroup) {
	r.POST("/shorter", Shorten)
	r.GET("/:shorter", Redirect)
}

func Shorten(c *gin.Context) {
	req := struct {
		URL    string `json:"url" binding:"required"`
		Day    int    `json:"day"`
		Hour   int    `json:"hour"`
		Minute int    `json:"minute"`
		Second int    `json:"second"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		errs.Fail(c, errs.InvalidRequest.WithOrigin(err))
		return
	}

	// deal with the time
	totalSeconds := req.Day*86400 + req.Hour*3600 + req.Minute*60 + req.Second
	expiredTime := time.Now().Add(time.Duration(totalSeconds) * time.Second)
	fmt.Println(expiredTime.Format("2006-01-02 15:04:05"))
	url := &model.Url{
		LongURL:  req.URL,
		ShortURL: shorter.Shorten(req.URL),
		Time:     expiredTime,
	}

	// deal with the short url
	err := pgsql.CreateUrl(url)
	for errors.Is(err, gorm.ErrDuplicatedKey) {
		url.ShortURL += shorter.Add1(1)
		err = pgsql.CreateUrl(url)
	}
	if err != nil {
		errs.Fail(c, errs.DatabaseError.WithOrigin(err))
		return
	}
	err = redis.CreateUrl(url)
	if err != nil {
		errs.Fail(c, errs.DatabaseError.WithOrigin(err))
		return
	}
	errs.Success(c, "/api/"+url.ShortURL)
}

func Redirect(c *gin.Context) {
	shorten := c.Param("shorter")
	url, _ := redis.GetUrl(shorten)
	if url != "" {
		fmt.Println("find record in redis success")
		c.Redirect(302, url)
		return
	}
	url, err := pgsql.GetUrl(shorten)
	if err != nil {
		errs.Fail(c, errs.DatabaseError.WithOrigin(err))
		return
	}
	fmt.Println("find record in pgsql success")
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}
	c.Redirect(302, url)
}
