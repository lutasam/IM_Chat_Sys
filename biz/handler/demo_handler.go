package handler

import (
	"github.com/gin-gonic/gin"
	"lutasam/GIN_LUTA/biz/bo"
	"lutasam/GIN_LUTA/biz/service"
	"lutasam/GIN_LUTA/biz/utils"
)

type DemoController struct{}

func RegisterDemoRouter(r *gin.Engine) {
	demoController := &DemoController{}
	Demo := r.Group("/demo")
	{
		Demo.GET("/ping", demoController.Ping)
		Demo.POST("/hello", demoController.Hello)
	}
}

func (ins *DemoController) Ping(c *gin.Context) {
	pong, err := service.GetDemoService().Ping(c)
	if err != nil {
		utils.Response(c, 400, "server error", nil)
		return
	}
	utils.Response(c, 200, "OK", pong)
}

func (ins *DemoController) Hello(c *gin.Context) {
	req := &bo.HelloRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.Response(c, 400, "server error", nil)
		return
	}
	hello, err := service.GetDemoService().Hello(c, req)
	if err != nil {
		utils.Response(c, 400, "server error", nil)
		return
	}
	utils.Response(c, 200, "OK", hello)
}
