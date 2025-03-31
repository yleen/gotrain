package initialize

import (
	"blog_backend/global"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"os"
)

func ConnectRedis() *redis.Client {
	redisCfg := global.Config.Redis

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Address,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		global.Log.Error("fail to connect redis", zap.Error(err))
		os.Exit(1)
	}

	return redisClient
}
