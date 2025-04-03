package api

import (
	"blog_backend/model/response"
	"github.com/gin-gonic/gin"
)

type baseApi struct{}

func (baseApi *baseApi) Captcha(c *gin.Context) {
	response.OKWithData("", c)
}
