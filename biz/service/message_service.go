package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/bo"
	"github.com/lutasam/chat/biz/common"
	"github.com/lutasam/chat/biz/dal"
	"github.com/lutasam/chat/biz/model"
	"github.com/lutasam/chat/biz/vo"
	"sync"
)

type MessageService struct{}

var (
	messageService     *MessageService
	messageServiceOnce sync.Once
)

func GetMessageService() *MessageService {
	messageServiceOnce.Do(func() {
		messageService = &MessageService{}
	})
	return messageService
}

func (ins *MessageService) GetUserMessages(c *gin.Context, sendID, receiveID uint64) (*bo.GetUserMessagesResponse, error) {
	user, err := dal.GetUserDal().GetUserByID(c, sendID)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, common.USERDOESNOTEXIST
	}
	user, err = dal.GetUserDal().GetUserByID(c, receiveID)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, common.USERDOESNOTEXIST
	}
	var messages []*model.UserMessage
	messages, err = dal.GetMessageDal().GetUserMessages(c, sendID, receiveID)
	if err != nil {
		return nil, err
	}
	messageVOs := convert2UserMessagesVO(messages)
	return &bo.GetUserMessagesResponse{
		Messages: messageVOs,
	}, nil
}

func convert2UserMessagesVO(messages []*model.UserMessage) []*vo.UserMessagesVO {
	var messageVOs []*vo.UserMessagesVO
	for _, message := range messages {
		messageVOs = append(messageVOs, &vo.UserMessagesVO{
			Name:    message.Name,
			Avatar:  message.Avatar,
			Content: message.Content,
		})
	}
	return messageVOs
}

func (ins *MessageService) GetGroupMessages(c *gin.Context, groupID uint64) (*bo.GetGroupMessagesResponse, error) {

}

func (ins *MessageService) GetUserMessageNum(c *gin.Context, receiveID uint64) (*bo.GetUserMessageNumResponse, error) {

}

func (ins *MessageService) GetGroupMessageNum(c *gin.Context, groupID uint64) (*bo.GetGroupMessageNumResponse, error) {

}

func (ins *MessageService) SendUserMessage(c *gin.Context, req *bo.SendUserMessageRequest) error {

}

func (ins *MessageService) SendGroupMessage(c *gin.Context, req *bo.SendGroupMessageRequest) error {

}
