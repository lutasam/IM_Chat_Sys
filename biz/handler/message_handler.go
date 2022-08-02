package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/bo"
	"github.com/lutasam/chat/biz/common"
	"github.com/lutasam/chat/biz/service"
	"github.com/lutasam/chat/biz/utils"
)

type MessageController struct{}

func RegisterMessageRouter(r *gin.RouterGroup) {
	messageController := &MessageController{}
	{
		r.GET("/get_user_messages/:receive_user_id", messageController.GetUserMessages)
		r.GET("/get_group_messages/:group_id", messageController.GetGroupMessages)
		r.GET("/get_user_message_num/:receive_user_id", messageController.GetUserMessageNum)
		r.GET("/get_group_message_num/:group_id", messageController.GetGroupMessageNum)
		r.POST("/send_message_to_user", messageController.SendUserMessage)
		r.POST("/send_message_to_group", messageController.SendGroupMessage)
	}
}

func (ins *MessageController) GetUserMessages(c *gin.Context) {
	jwtStruct, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		utils.ResponseClientError(c, common.USERNOTLOGIN)
		return
	}
	var receiveID uint64
	receiveID, err = utils.ParseString2Uint64(c.Param("receive_user_id"))
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	var resp *bo.GetUserMessagesResponse
	resp, err = service.GetMessageService().GetUserMessages(c, jwtStruct.UserID, receiveID)
	if err != nil {
		if errors.Is(err, common.USERDOESNOTEXIST) {
			utils.ResponseClientError(c, common.USERDOESNOTEXIST)
			return
		} else {
			utils.ResponseServerError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *MessageController) GetGroupMessages(c *gin.Context) {
	groupID, err := utils.ParseString2Uint64(c.Param("group_id"))
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	var resp *bo.GetGroupMessagesResponse
	resp, err = service.GetMessageService().GetGroupMessages(c, groupID)
	if err != nil {
		if errors.Is(err, common.USERDOESNOTEXIST) {
			utils.ResponseClientError(c, common.USERDOESNOTEXIST)
			return
		} else if errors.Is(err, common.GROUPNOTEXIST) {
			utils.ResponseClientError(c, common.GROUPNOTEXIST)
			return
		} else {
			utils.ResponseServerError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *MessageController) GetUserMessageNum(c *gin.Context) {
	jwtStruct, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		utils.ResponseClientError(c, common.USERNOTLOGIN)
		return
	}
	var receiveID uint64
	receiveID, err = utils.ParseString2Uint64(c.Param("receive_user_id"))
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	var resp *bo.GetUserMessageNumResponse
	resp, err = service.GetMessageService().GetUserMessageNum(c, jwtStruct.UserID, receiveID)
	if err != nil {
		if errors.Is(err, common.USERDOESNOTEXIST) {
			utils.ResponseClientError(c, common.USERDOESNOTEXIST)
			return
		} else {
			utils.ResponseServerError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *MessageController) GetGroupMessageNum(c *gin.Context) {
	groupID, err := utils.ParseString2Uint64(c.Param("group_id"))
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	var resp *bo.GetGroupMessageNumResponse
	resp, err = service.GetMessageService().GetGroupMessageNum(c, groupID)
	if err != nil {
		if errors.Is(err, common.USERDOESNOTEXIST) {
			utils.ResponseClientError(c, common.USERDOESNOTEXIST)
			return
		} else if errors.Is(err, common.GROUPNOTEXIST) {
			utils.ResponseClientError(c, common.GROUPNOTEXIST)
			return
		} else {
			utils.ResponseServerError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *MessageController) SendUserMessage(c *gin.Context) {
	jwtStruct, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		utils.ResponseClientError(c, common.USERNOTLOGIN)
		return
	}
	req := &bo.SendUserMessageRequest{}
	err = c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	err = service.GetMessageService().SendUserMessage(c, req, jwtStruct.UserID)
	if err != nil {
		if errors.Is(err, common.USERDOESNOTEXIST) {
			utils.ResponseClientError(c, common.USERDOESNOTEXIST)
			return
		} else {
			utils.ResponseServerError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, nil)
}

func (ins *MessageController) SendGroupMessage(c *gin.Context) {
	jwtStruct, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		utils.ResponseClientError(c, common.USERNOTLOGIN)
		return
	}
	req := &bo.SendGroupMessageRequest{}
	err = c.ShouldBind(req)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	err = service.GetMessageService().SendGroupMessage(c, req, jwtStruct.UserID)
	if err != nil {
		if errors.Is(err, common.USERDOESNOTEXIST) {
			utils.ResponseClientError(c, common.USERDOESNOTEXIST)
			return
		} else if errors.Is(err, common.GROUPNOTEXIST) {
			utils.ResponseClientError(c, common.GROUPNOTEXIST)
			return
		} else {
			utils.ResponseServerError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, nil)
}
