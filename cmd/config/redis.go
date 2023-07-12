package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	Redis *redis.Client
)

func ConnectRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, // use default DB
	})
	if err := Redis.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}
}
