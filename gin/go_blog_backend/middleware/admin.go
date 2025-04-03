package middleware

import (
	"blog_backend/model/appTypes"
	"blog_backend/model/response"
	"blog_backend/utils"
	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID := utils.GetRouleID(c)

		if roleID != appTypes.Admin {
			response.Forbidden("Access denied. Admin privileges are required", c)
			c.Abort()
			return
		}

		c.Next()
	}
}
