package global

import (
	"blog_backend/config"
	"go.uber.org/zap"
)

var (
	Config *config.Config
	Log    *zap.Logger
)
