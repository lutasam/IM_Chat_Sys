package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/bo"
	"github.com/lutasam/chat/biz/common"
	"github.com/lutasam/chat/biz/service"
	"github.com/lutasam/chat/biz/utils"
)

type GroupController struct{}

func RegisterGroupRouter(r *gin.RouterGroup) {
	groupController := &GroupController{}
	{
		r.POST("/create_group", groupController.CreateGroup)
		r.POST("/update_group", groupController.UpdateGroup)
		r.GET("/get_group_detail/:group_id", groupController.GetGroupDetail)
		r.GET("/get_all_groups", groupController.GetAllGroups)
	}
}

func (ins *GroupController) CreateGroup(c *gin.Context) {
	jwtStruct, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		utils.ResponseClientError(c, common.USERNOTLOGIN)
		return
	}
	req := &bo.CreateGroupRequest{}
	err = c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	err = service.GetGroupService().CreateGroup(c, req, jwtStruct.UserID)
	if err != nil {
		if errors.Is(err, common.USERNOTLOGIN) {
			utils.ResponseClientError(c, common.USERNOTLOGIN)
			return
		} else {
			utils.ResponseServerError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, nil)
}

func (ins *GroupController) UpdateGroup(c *gin.Context) {

}

func (ins *GroupController) GetGroupDetail(c *gin.Context) {

}

func (ins *GroupController) GetAllGroups(c *gin.Context) {

}
