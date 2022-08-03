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

func RegisterLoginRouter(r *gin.RouterGroup) {
	loginController := &LoginController{}
	{
		r.POST("/do_login", loginController.DoLogin)
		r.POST("/do_register", loginController.DoRegister)
	}
}

func (ins *LoginController) DoLogin(c *gin.Context) {
	req := &bo.LoginRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetLoginService().DoLogin(c, req)
	if err != nil {
		if errors.Is(err, common.USERDOESNOTEXIST) {
			utils.ResponseClientError(c, common.USERDOESNOTEXIST)
			return
		} else if errors.Is(err, common.PASSWORDISERROR) {
			utils.ResponseClientError(c, common.PASSWORDISERROR)
			return
		} else if errors.Is(err, common.USERINPUTERROR) {
			utils.ResponseClientError(c, common.USERINPUTERROR)
			return
		} else {
			utils.ResponseServerError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *LoginController) DoRegister(c *gin.Context) {
	req := &bo.RegisterRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetLoginService().DoRegister(c, req)
	if err != nil {
		if errors.Is(err, common.USEREXISTED) {
			utils.ResponseClientError(c, common.USEREXISTED)
			return
		} else if errors.Is(err, common.USERINPUTERROR) {
			utils.ResponseClientError(c, common.USERINPUTERROR)
			return
		} else {
			utils.ResponseServerError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}
