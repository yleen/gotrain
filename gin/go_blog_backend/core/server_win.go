//go:build windows
// +build windows

package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func initServer(addr string, router *gin.Engine) server {
	return &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
