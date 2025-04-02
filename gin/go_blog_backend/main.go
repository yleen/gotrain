package main

import (
	"blog_backend/core"
	"blog_backend/global"
	"blog_backend/initialize"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func main() {
	global.Config = core.InitConf()
	global.Log = core.InitLogger()
	global.DB = initialize.InitGorm()
	global.Redis = initialize.ConnectRedis() // diff
	global.ESClient = initialize.ConnectEs()

	defer func(Redis *redis.Client) {
		err := Redis.Close()
		if err != nil {
			global.Log.Error("Error closing redis connection", zap.Error(err))
		}
	}(global.Redis)

	//flag.InitFlag()

	initialize.InitCron()

	core.RunServer()
}
