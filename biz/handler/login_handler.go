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
	var resp *bo.LoginResponse
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseError(c, common.USERINPUTERROR)
		return
	}
	resp, err = service.GetLoginService().DoLogin(c, req)
	if err != nil {
		if errors.Is(err, common.USERDOESNOTEXIST) {
			utils.ResponseError(c, common.USERDOESNOTEXIST)
			return
		} else if errors.Is(err, common.PASSWORDISERROR) {
			utils.ResponseError(c, common.PASSWORDISERROR)
			return
		} else if errors.Is(err, common.USERINPUTERROR) {
			utils.ResponseError(c, common.USERINPUTERROR)
			return
		} else {
			utils.ResponseError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *LoginController) DoRegister(c *gin.Context) {
	req := &bo.RegisterRequest{}
	var resp *bo.RegisterResponse
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseError(c, common.USERINPUTERROR)
		return
	}
	resp, err = service.GetLoginService().DoRegister(c, req)
	if err != nil {
		if errors.Is(err, common.USEREXISTED) {
			utils.ResponseError(c, common.USEREXISTED)
			return
		} else if errors.Is(err, common.USERINPUTERROR) {
			utils.ResponseError(c, common.USERINPUTERROR)
			return
		} else {
			utils.ResponseError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}
