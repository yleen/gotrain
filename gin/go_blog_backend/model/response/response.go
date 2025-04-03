package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func OKWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "success", c)
}

func NoAuth(message string, c *gin.Context) {
	Result(ERROR, gin.H{"reload": true}, message, c)
}

func Forbidden(message string, c *gin.Context) {
	c.JSON(http.StatusForbidden, Response{
		Code: ERROR,
		Data: nil,
		Msg:  message,
	})
}
