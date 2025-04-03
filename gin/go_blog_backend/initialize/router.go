package initialize

import (
	"blog_backend/global"
	"blog_backend/middleware"
	"blog_backend/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	var store = cookie.NewStore([]byte(global.Config.System.SessionsSecret))
	router.Use(sessions.Sessions("session", store))
	router.StaticFS(global.Config.Upload.Path, http.Dir(global.Config.Upload.Path))
	routerGroup := routes.RouterGroupApp

	publicGroup := router.Group(global.Config.System.RouterPrefix)
	privateGroup := router.Group(global.Config.System.RouterPrefix)
	privateGroup.Use(middleware.JWTAuth())
	adminGroup := router.Group(global.Config.System.RouterPrefix)
	adminGroup.Use(middleware.JWTAuth()).Use(middleware.AdminAuth())
	{
		routerGroup.BaseRouter.InitBaseRouter(publicGroup)
	}
	return router
}
