package utils

import (
	"blog_backend/global"
	"blog_backend/model/appTypes"
	"blog_backend/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

func GetClaims(c *gin.Context) (*request.JwtCustomClaims, error) {
	token := GetAccessToken(c)
	j := NewJWT()
	claims, err := j.ParseAccessToken(token)
	if err != nil {
		global.Log.Error("Failed to retrieve JWT parsing information from Gin's Context. Please check if the request header contains 'x-access-token' or if the claims structure is correct", zap.Error(err))
	}
	return claims, nil
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

func GetRouleID(c *gin.Context) appTypes.RoleID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return appTypes.Guest
		} else {
			return cl.RoleID
		}
	} else {
		waitUse := claims.(*request.JwtCustomClaims)
		return waitUse.RoleID
	}
}
