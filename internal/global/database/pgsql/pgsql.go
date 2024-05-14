package pgsql

import (
	"URL_shorter/internal/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=127.0.0.1 user=gorm password=gorm dbname=shorter port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		fmt.Println("postgresql init failed")
	}
	err = db.AutoMigrate(
		model.Url{},
	)
	if err != nil {
		panic(err)
	}
	DB = db
}
