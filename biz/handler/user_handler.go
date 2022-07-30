package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/bo"
	"github.com/lutasam/chat/biz/common"
	"github.com/lutasam/chat/biz/service"
	"github.com/lutasam/chat/biz/utils"
)

type UserController struct{}

func RegisterUserRouter(r *gin.RouterGroup) {
	userController := &UserController{}
	{
		r.GET("/get_detail/", userController.GetDetail)
		r.POST("/update_user_info", userController.UpdateUserInfo)
	}
}

func (ins *UserController) GetDetail(c *gin.Context) {
	jwtStruct, exist := c.Get("jwtStruct")
	if !exist {
		utils.ResponseError(c, common.USERNOTLOGIN)
		return
	}
	resp, err := service.GetUserService().GetUserDetail(c, jwtStruct.(*utils.JWTStruct).UserID)
	if err != nil {
		if errors.Is(err, common.USERDOESNOTEXIST) {
			utils.ResponseError(c, common.USERDOESNOTEXIST)
			return
		} else {
			utils.ResponseError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *UserController) UpdateUserInfo(c *gin.Context) {
	req := &bo.UpdateUserInfoRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseError(c, common.USERINPUTERROR)
		return
	}
	err = service.GetUserService().UpdateUserInfo(c, req)
	if err != nil {
		if errors.Is(err, common.USERDOESNOTEXIST) {
			utils.ResponseError(c, common.USERDOESNOTEXIST)
			return
		} else {
			utils.ResponseError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, nil)
}
