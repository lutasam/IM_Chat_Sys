package handler

import (
	"github.com/gin-gonic/gin"
)

type MessageController struct{}

func RegisterMessageRouter(r *gin.RouterGroup) {
	messageController := &MessageController{}
	{
		r.POST("/get_conversation_records", messageController.GetConversationRecords)
		r.POST("/get_group_conversation_records", messageController.GetGroupConversationRecords)
		r.GET("get_message_num", messageController.GetMessageNum)
	}
}

func (ins *MessageController) GetConversationRecords(c *gin.Context) {

}

func (ins *MessageController) GetGroupConversationRecords(c *gin.Context) {

}

func (ins *MessageController) GetMessageNum(c *gin.Context) {

}
