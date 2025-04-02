package middleware

import (
	"blog_backend/model/response"
	"blog_backend/service"
	"blog_backend/utils"
	"github.com/gin-gonic/gin"
)

var jwtService = service.ServiceGroupApp.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := utils.GetAccessToken(c)
		refreshToken := utils.GetRefreshToken(c)

		if jwtService.IsInBlackList(accessToken) {
			utils.ClearRefreshToken(c)
			response.NoAuth("Account logged in from another location or token is invalid", c)
			c.Abort()
			return
		}

	}
}
