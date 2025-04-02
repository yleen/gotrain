package core

import (
	"blog_backend/global"
	"blog_backend/initialize"
	"blog_backend/service"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	addr := global.Config.System.Addr()

	router := initialize.InitRouter()

	service.LoadAll()

	s := initServer(addr, router)
	global.Log.Info("Server is running on ", zap.String("address", addr))
	global.Log.Error(s.ListenAndServe().Error())
}
