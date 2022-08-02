package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/utils"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	InitRouterAndMiddleware(r)

	err := r.Run(":" + utils.GetConfigResolve().GetConfigString("server.port"))
	if err != nil {
		panic(err)
	}
}
