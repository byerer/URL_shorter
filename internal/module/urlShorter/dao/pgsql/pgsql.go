package pgsql

import (
	db "URL_shorter/internal/global/database/pgsql"
	"URL_shorter/internal/model"
	"errors"
	"time"
)

func CreateUrl(url *model.Url) error {
	err := db.DB.Create(url).Error
	return err
}

func GetUrl(shorten string) (string, error) {
	var record model.Url
	err := db.DB.Where("short_url = ?", shorten).First(&record).Error
	if time.Now().After(record.Time) {
		return "", errors.New("url expired")
	}
	return record.LongURL, err
}
