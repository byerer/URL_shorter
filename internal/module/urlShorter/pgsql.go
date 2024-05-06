package urlShorter

import (
	db "URL_shorter/internal/global/database"
	"URL_shorter/internal/model"
)

func CreateUrl(url *model.Url) error {
	err := db.DB.Create(url).Error
	return err
}

func GetUrl(shorten string) (string, error) {
	var record model.Url
	err := db.DB.Where("short_url = ?", shorten).First(&record).Error
	return record.LongURL, err
}
