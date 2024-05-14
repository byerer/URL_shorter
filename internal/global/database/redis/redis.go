package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func Init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	RDB = rdb
}

func Close() {
	RDB.Close()
}

func Ping() {
	_, err := RDB.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
