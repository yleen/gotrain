package initialize

import (
	"blog_backend/global"
	"blog_backend/utils"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"os"
)

func OtherInit() {
	refreshTokenExpiry, err := utils.ParseDuration(global.Config.Jwt.RefreshTokenExpireTime)
	if err != nil {
		global.Log.Error("Failed to parse refresh token expire time", zap.Error(err))
		os.Exit(1)
	}

	_, err = utils.ParseDuration(global.Config.Jwt.AccessTokenExpireTime)
	if err != nil {
		global.Log.Error("Failed to parse access token expire time", zap.Error(err))
		os.Exit(1)
	}

	global.BlackCache = local_cache.NewCache(local_cache.SetDefaultExpire(refreshTokenExpiry))
}
