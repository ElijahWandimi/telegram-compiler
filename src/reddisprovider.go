package src

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/oyamo/telegram-compiler/config"
)


func RedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_ENDPOINT,
		Password: config.REDIS_PASSWORD, // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
    fmt.Println(pong, err)
	return client, err
}