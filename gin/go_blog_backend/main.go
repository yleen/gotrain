package main

import (
	"blog_backend/core"
	"blog_backend/global"
)

func main() {
	global.Config = core.InitConf()
	global.Log = core.InitLogger()
}
