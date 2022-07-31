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
		r.POST("/send_message_to_user", messageController.SendMessageToUser)
		r.POST("/send_message_to_group", messageController.SendMessageToGroup)
	}
}

func (ins *MessageController) GetUserMessages(c *gin.Context) {
	jwtStruct, exist := c.Get("jwtStruct")
	if !exist {
		utils.ResponseError(c, common.USERNOTLOGIN)
		return
	}
	receive_id, err := utils.ParseString2Uint64(c.Param("receive_user_id"))
	if err != nil {
		utils.ResponseError(c, common.USERINPUTERROR)
		return
	}
	var resp *bo.GetUserMessagesResponse
	resp, err = service.GetMessageService().GetUserMessages(c, jwtStruct.(utils.JWTStruct).UserID, receive_id)
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

func (ins *MessageController) GetGroupMessages(c *gin.Context) {

}

func (ins *MessageController) GetUserMessageNum(c *gin.Context) {

}

func (ins *MessageController) GetGroupMessageNum(c *gin.Context) {

}

func (ins *MessageController) SendMessageToUser(c *gin.Context) {

}

func (ins *MessageController) SendMessageToGroup(c *gin.Context) {

}
