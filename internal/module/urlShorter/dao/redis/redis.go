package redis

import (
	"URL_shorter/internal/global/database/redis"
	"URL_shorter/internal/model"
	"time"
)

func CreateUrl(url *model.Url) error {
	err := redis.Rdb.Set(redis.Ctx, url.ShortURL, url.LongURL, 2*url.Time.Sub(time.Now())/10).Err()
	return err
}

func GetUrl(shorten string) (string, error) {
	longURL, err := redis.Rdb.Get(redis.Ctx, shorten).Result()
	if err != nil {
		return "", err
	}
	return longURL, nil
}
