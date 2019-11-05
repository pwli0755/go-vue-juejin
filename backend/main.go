package main

import (
	"backend/conf"
	"backend/server"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	if !conf.Conf.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	// 装载路由
	r := server.NewRouter()

	r.Run(fmt.Sprintf(":%s", conf.Conf.Server.Port))
}
