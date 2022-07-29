package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/bo"
	"github.com/lutasam/chat/biz/common"
	"github.com/lutasam/chat/biz/service"
	"github.com/lutasam/chat/biz/utils"
)

type LoginController struct{}

func RegisterLoginRouter(r *gin.Engine) {
	loginController := &DemoController{}
	login := r.Group("/login")
	{
		login.GET("/do_login", loginController.Ping)
		login.POST("/do_register", loginController.Hello)
	}
}

func (ins *LoginController) DoLogin(c *gin.Context) {
	req := &bo.LoginRequest{}
	var resp *bo.LoginResponse
	err := c.ShouldBind(req)
	if err != nil {
		utils.Response(c, common.USERINPUTERROR.Code(), common.USERINPUTERROR.Error(), nil)
		return
	}
	resp, err = service.GetLoginService().DoLogin(c, req)
	if err != nil {
		if errors.Is(err, common.USERNAMEDOESNOTEXIST) {
			utils.Response(c, common.USERNAMEDOESNOTEXIST.Code(), common.USERNAMEDOESNOTEXIST.Error(), nil)
			return
		} else if errors.Is(err, common.PASSWORDISERROR) {
			utils.Response(c, common.PASSWORDISERROR.Code(), common.PASSWORDISERROR.Error(), nil)
			return
		}
	}
	utils.Response(c, common.STATUSOKCODE, common.STATUSOKMSG, resp)
}

func (ins *LoginController) DoResigter(c *gin.Context) {

}

func (ins *LoginController) Ping(c *gin.Context) {
	pong, err := service.GetDemoService().Ping(c)
	if err != nil {
		utils.Response(c, 400, "server error", nil)
		return
	}
	utils.Response(c, 200, "OK", pong)
}

func (ins *LoginController) Hello(c *gin.Context) {
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
