package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var (
	address = "localhost:6379" //move to os.env
)

type redisClient struct {
	RedisDb *redis.Client
}

func newRedisClient() *redisClient {
	client := &redisClient{
		RedisDb: redis.NewClient(&redis.Options{
			Addr:     address,
			Password: "",
			DB:       0,
		}),
	}
	fmt.Printf("\nConnected to Redis at: %v\n\n", client.RedisDb.Options().Addr)
	return client
}

func (rc *redisClient) SetKeyValue(key string, value string) {

	if err := rc.RedisDb.Set(ctx, key, value, 0).Err(); err == nil {
		fmt.Printf("\nVALUE: %v set for KEY: %v\n\n", value, key)
	} else {
		panic(err)
	}
}

func (rc *redisClient) GetKeyValue(key string) {

	val, err := rc.RedisDb.Get(ctx, key).Result()

	if err == redis.Nil {
		fmt.Printf("\nKEY: %v does not exist\n\n", key)
	} else if err != nil {
		panic(err)
	} else {
		fmt.Printf("\nKEY: %v, VALUE: %v\n\n", key, val)
	}
}
