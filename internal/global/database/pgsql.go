package database

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

func Ping() {
	var result int
	// 使用 Raw SQL 查询并扫描结果
	err := DB.Raw("SELECT 1+1").Scan(&result).Error
	if err != nil {
		fmt.Println("Failed to ping PostgreSQL:", err)
		return
	}
	fmt.Println("PostgreSQL ping success, answer is", result)
}
