package redis

import (
	"fmt"

	"github.com/go-redis/redis/v7"
	"github.com/zeelrupapara/chat/config"
)

var client *redis.Client

func InitRedis(conf config.AppConfig) {
	client = redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		DB:       0,
		Password: conf.Redis.Password,
	})

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Error connecting to redis:", err)
	} else {
		fmt.Println("Redis Connected")
	}
}

func SaveMessage(channel, message string) {
	err := client.Publish(channel, message).Err()
	if err != nil {
		fmt.Println("Error pub message:", err)
	}
}