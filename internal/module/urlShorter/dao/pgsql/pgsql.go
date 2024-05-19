package pgsql

import (
	db "URL_shorter/internal/global/database/pgsql"
	"URL_shorter/internal/model"
	"errors"
	"fmt"
	"time"
)

func CreateUrl(url *model.Url) error {
	err := db.DB.Create(url).Error
	fmt.Println(url.Time.Format("2006-01-02 15:04:05"))
	return err
}

func GetUrl(shorten string) (string, error) {
	var record model.Url
	err := db.DB.Where("short_url = ?", shorten).First(&record).Error
	fmt.Println(record.Time.Format("2006-01-02 15:04:05"))
	if time.Now().After(record.Time) {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), record.Time.Format("2006-01-02 15:04:05"))
		return "", errors.New("url expired")
	}
	return record.LongURL, err
}
