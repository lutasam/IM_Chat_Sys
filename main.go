package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/utils"
)

func main() {
	r := gin.New()

	gin.SetMode(gin.ReleaseMode)

	InitRouterAndMiddleware(r)

	err := r.Run(":" + utils.GetConfigResolve().GetConfigString("server.port"))
	if err != nil {
		panic(err)
	}
}
