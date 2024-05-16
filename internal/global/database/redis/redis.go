package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client
var Ctx = context.Background()

func Init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	Rdb = rdb
}

func Close() {
	Rdb.Close()
}

func Ping() {
	_, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
