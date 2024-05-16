package model

import (
	"gorm.io/gorm"
	"time"
)

type Url struct {
	gorm.Model
	ShortURL string `gorm:"unique;type:varchar(255);not null;"`
	LongURL  string `gorm:"type:varchar(255);not null;"`
	Time     time.Time
}
