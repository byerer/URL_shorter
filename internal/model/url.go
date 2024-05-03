package model

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	ShortURL string `gorm:"unique;type:varchar(255);not null;"`
	LongURL  string `gorm:"type:varchar(255);not null;"`
}
