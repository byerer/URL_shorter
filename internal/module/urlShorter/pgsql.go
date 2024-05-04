package urlShorter

import (
	db "URL_shorter/internal/global/database"
	"URL_shorter/internal/model"
	st "URL_shorter/service/shorten"
)

func CreateUrl(url string) error {
	short := st.Shorten(url)
	record := model.Url{ShortURL: short, LongURL: url}
	err := db.DB.Create(&record).Error
	return err
}

func GetUrl(shorten string) (string, error) {
	var record model.Url
	err := db.DB.Where("short_url = ?", shorten).First(&record).Error
	return record.LongURL, err
}
