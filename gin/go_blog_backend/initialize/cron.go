package initialize

import (
	"blog_backend/global"
	"blog_backend/task"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"os"
)

type ZapLogger struct {
	logger *zap.Logger
}

func (z ZapLogger) Info(msg string, keysAndValues ...interface{}) {
	z.logger.Info(msg, zap.Any("KeysAndValues", keysAndValues))
}

func (z ZapLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	z.logger.Error(msg, zap.Any("KeysAndValues", keysAndValues), zap.Error(err))
}

func NewZapLogger() *ZapLogger {
	return &ZapLogger{logger: global.Log}
}

func InitCron() {
	c := cron.New(cron.WithLogger(NewZapLogger()))
	err := task.RegisterScheduledTasks(c)
	if err != nil {
		global.Log.Error("Error scheduling cron task", zap.Error(err))
		os.Exit(1)
	}
	c.Start()
}
