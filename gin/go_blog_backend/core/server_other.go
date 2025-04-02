//go:build !windows

package core

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

func initServer(adress string, router *gin.Engine) server {
	s := endless.NewServer(adress, router)
	s.ReadHeaderTimeout = time.Second * 10
	s.WriteTimeout = time.Second * 10
	s.MaxHeaderBytes = 1 << 20

	return s
}
