package routes

import (
	"blog_backend/api"
	"github.com/gin-gonic/gin"
)

type baseRouter struct{}

func (b *baseRouter) InitBaseRouter(router *gin.RouterGroup) {
	baseRouter := router.Group("base")

	baseApi := api.ApiGroupApp.BaseApi

	baseRouter.POST("captcha", baseApi.Captcha)
}
