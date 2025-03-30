package config

import (
	"exchangeapp_train/global"
	"log"

	"github.com/go-redis/redis"
)

func initRedis() {
	addr := Appconfig.Redis.Addr
	db := Appconfig.Redis.DB
	password := Appconfig.Redis.Password

	RedisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       db,
		Password: password,
	})

	_, err := RedisClient.Ping().Result()

	if err != nil {
		log.Fatalf("Failed to connect to Redis, got error: %v", err)
	}

	global.RedisDB = RedisClient
}
