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
		r.GET("/find_groups/:input_str", groupController.FindGroups)
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
	req := &bo.UpdateGroupRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	err = service.GetGroupService().UpdateGroup(c, req)
	if err != nil {
		if errors.Is(err, common.GROUPNOTEXIST) {
			utils.ResponseClientError(c, common.GROUPNOTEXIST)
			return
		} else if errors.Is(err, common.USERDOESNOTEXIST) {
			utils.ResponseClientError(c, common.USERDOESNOTEXIST)
			return
		} else {
			utils.ResponseServerError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, nil)
}

func (ins *GroupController) GetGroupDetail(c *gin.Context) {
	groupIDStr := c.Param("group_id")
	groupID, err := utils.ParseString2Uint64(groupIDStr)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetGroupService().GetGroupDetail(c, groupID)
	if err != nil {
		if errors.Is(err, common.GROUPNOTEXIST) {
			utils.ResponseClientError(c, common.GROUPNOTEXIST)
			return
		} else if errors.Is(err, common.USERDOESNOTEXIST) {
			utils.ResponseClientError(c, common.USERDOESNOTEXIST)
			return
		} else {
			utils.ResponseServerError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *GroupController) GetAllGroups(c *gin.Context) {
	jwtStruct, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		utils.ResponseClientError(c, common.USERNOTLOGIN)
		return
	}
	resp, err := service.GetGroupService().GetAllGroups(c, jwtStruct.UserID)
	if err != nil {
		if errors.Is(err, common.USERNOTLOGIN) {
			utils.ResponseClientError(c, common.USERNOTLOGIN)
			return
		} else if errors.Is(err, common.USERDOESNOTEXIST) {
			utils.ResponseClientError(c, common.USERDOESNOTEXIST)
			return
		} else {
			utils.ResponseServerError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *GroupController) FindGroups(c *gin.Context) {
	str := c.Param("input_str")
	resp, err := service.GetGroupService().FindGroups(c, str)
	if err != nil {
		if errors.Is(err, common.USERINPUTERROR) {
			utils.ResponseClientError(c, common.USERINPUTERROR)
			return
		} else {
			utils.ResponseServerError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}
