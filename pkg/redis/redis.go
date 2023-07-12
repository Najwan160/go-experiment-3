package redis

import (
	"context"
	"fmt"

	"github.com/Najwan160/go-experiment-3/code/domain/base"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	redis *redis.Client
}

func NewRedis(redis *redis.Client) base.Redis {
	return &Redis{redis}
}

func (rd *Redis) SetKey(ctx context.Context, key string, value interface{}) {

	err := rd.redis.Set(ctx, key, value, 0).Err()
	if err != nil {
		fmt.Printf("unable to SET data. error: %v", err)
		return
	}
}

func (rd *Redis) GetKey(ctx context.Context, key string) (value string) {

	get := rd.redis.Get(ctx, key)
	if err := get.Err(); err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return
	}
	getResult, err := get.Result()
	if err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return
	}
	return getResult
}
