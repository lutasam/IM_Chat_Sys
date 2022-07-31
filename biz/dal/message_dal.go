package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/common"
	"github.com/lutasam/chat/biz/model"
	"github.com/lutasam/chat/biz/repository"
	"sync"
)

type MessageDal struct{}

var (
	messageDal     *MessageDal
	messageDalOnce sync.Once
)

func GetMessageDal() *MessageDal {
	messageDalOnce.Do(func() {
		messageDal = &MessageDal{}
	})
	return messageDal
}

func (ins *MessageDal) CreateUserMessage(c *gin.Context, message *model.UserMessage) error {
	err := repository.GetDB().Table(message.TableName()).Create(message).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *MessageDal) GetUserMessages(c *gin.Context, sendID, receiveID uint64) ([]*model.UserMessage, error) {
	var messages []*model.UserMessage
	err := repository.GetDB().Table(model.UserMessage{}.TableName()).Find(messages).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return messages, nil
}

func (ins *MessageDal) CreateGroupMessage(c *gin.Context, message *model.GroupMessage) error {
	err := repository.GetDB().Table(model.GroupMessage{}.TableName()).Create(message).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *MessageDal) GetGroupMessages(c *gin.Context, groupID uint64) ([]*model.GroupMessage, error) {
	var messages []*model.GroupMessage
	err := repository.GetDB().Table(model.GroupMessage{}.TableName()).Find(messages).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return messages, nil
}
