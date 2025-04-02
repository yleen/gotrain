package utils

import (
	"github.com/gin-gonic/gin"
	"net"
)

func GetAccessToken(c *gin.Context) string {
	token := c.Request.Header.Get("x-access-token")
	return token
}

func GetRefreshToken(c *gin.Context) string {
	token := c.Request.Header.Get("x-refresh-token")
	return token
}

func ClearRefreshToken(c *gin.Context) {
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	setCookie(c, "x-refresh-token", "", -1, host)
}

func setCookie(c *gin.Context, name, value string, maxAge int, host string) {
	if net.ParseIP(host) != nil {
		c.SetCookie(name, value, maxAge, "/", "", false, true)
	} else {
		c.SetCookie(name, value, maxAge, "/", host, false, true)
	}
}
