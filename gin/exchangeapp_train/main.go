package main

import (
	"exchangeapp_train/config"
	"exchangeapp_train/router"
)

func main() {
	config.InitConfig()

	r := router.SetupRouter()

	port := config.Appconfig.App.Port

	if port == "" {
		port = ":8080"
	}

	r.Run(port)
}
