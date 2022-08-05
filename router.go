package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/handler"
	"github.com/lutasam/chat/biz/middleware"
)

// InitRouterAndMiddleware init methods, you can register middleware and router here
func InitRouterAndMiddleware(r *gin.Engine) {
	// TODO: this config is gin.Logger's config, do not use because we use our logger system
	// set log output filepath
	// logFile, err := os.Create(utils.GetConfigResolve().GetConfigString("log.filepath"))
	// if err != nil {
	// 	panic(err)
	// }
	// gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	// register global middlewares Logger and Recovery
	// r.Use(gin.Logger()) // gin's default Logger
	r.Use(middleware.RequestDoTracerId()) // our Logger
	r.Use(gin.Recovery())

	// register modules
	// login&&register module
	login := r.Group("/login")
	handler.RegisterLoginRouter(login)

	// friend module
	friend := r.Group("/friend")
	friend.Use(middleware.JWTAuth())
	handler.RegisterFriendRouter(friend)

	// message module
	message := r.Group("/message")
	message.Use(middleware.JWTAuth())
	handler.RegisterMessageRouter(message)

	// user module
	user := r.Group("/user")
	user.Use(middleware.JWTAuth())
	handler.RegisterUserRouter(user)

	// group module
	group := r.Group("/group")
	group.Use(middleware.JWTAuth())
	handler.RegisterGroupRouter(group)
}
