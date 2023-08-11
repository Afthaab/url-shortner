package db

import (
	"context"
	"os"

	"github.com/go-redis/redis"
)

var Ctx = context.Background()

func CreateRedisClientZero() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASSWORD"), // no password set
		DB:       0,                        // use default DB 0
	})
	return rdb
}

func CreateRedisClientOne() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Password: os.Getenv("DB_PASSWORD"), // no password set
		DB:       1,                        // use default DB 1
	})
	return rdb
}
