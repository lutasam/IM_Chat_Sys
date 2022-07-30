package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/handler"
	"github.com/lutasam/chat/biz/middleware"
	"github.com/lutasam/chat/biz/utils"
	"io"
	"os"
)

func InitRouterAndMiddleware(r *gin.Engine) {
	// 设置log文件输出
	logFile, err := os.Create(utils.GetConfigResolve().GetConfigString("log.filepath"))
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	// 注册全局中间件Recovery和Logger
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 注册分组路由
	// 登录、注册模块
	login := r.Group("/login")
	handler.RegisterLoginRouter(login)

	// 好友模块
	friend := r.Group("/friend")
	friend.Use(middleware.JWTAuth())

	// 对话模块
	message := r.Group("/message")
	message.Use(middleware.JWTAuth())

	// 用户模块
	user := r.Group("/user")
	user.Use(middleware.JWTAuth())
	handler.RegisterUserRouter(user)
}
