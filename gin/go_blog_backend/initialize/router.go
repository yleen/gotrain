package initialize

import (
	"blog_backend/global"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.New()
	router.Use()
	return router
}
