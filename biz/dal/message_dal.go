package dal

import (
	"gorm.io/gorm"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/common"
	"github.com/lutasam/chat/biz/model"
	"github.com/lutasam/chat/biz/repository"
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
	err := repository.GetDB().Table(model.UserMessage{}.TableName()).
		Where("send_user_id = ? and receive_user_id = ?", sendID, receiveID).Find(&messages).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if len(messages) == 0 {
		return nil, common.DATANOTFOUND
	}
	return messages, nil
}

func (ins *MessageDal) GetUserMessageNum(c *gin.Context, sendID, receiveID uint64) (int64, error) {
	var cnt *int64
	err := repository.GetDB().Table(model.UserMessage{}.TableName()).
		Where("send_user_id = ? and receive_user_id = ? and is_read = ?", sendID, receiveID, 0).Count(cnt).Error
	if err != nil {
		return 0, common.DATABASEERROR
	}
	return *cnt, nil
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
	err := repository.GetDB().Table(model.GroupMessage{}.TableName()).
		Where("group_id = ?", groupID).Find(&messages).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if len(messages) == 0 {
		return nil, common.DATANOTFOUND
	}
	return messages, nil
}

func (ins *MessageDal) GetGroupMessageNum(c *gin.Context, groupID uint64) (int64, error) {
	var cnt *int64
	err := repository.GetDB().Table(model.GroupMessage{}.TableName()).
		Where("group_id = ? and is_read = ?", groupID, 0).Find(cnt).Error
	if err != nil {
		return 0, common.DATABASEERROR
	}
	return *cnt, nil
}

func (ins *MessageDal) UpdateUserMessages(c *gin.Context, sendID, receiveID uint64) error {
	err := repository.GetDB().Session(&gorm.Session{AllowGlobalUpdate: true}).Table(model.UserMessage{}.TableName()).
		Where("send_user_id = ? and receive_user_id = ?", sendID, receiveID).Update("is_read", 1).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *MessageDal) UpdateGroupMessages(c *gin.Context, groupID uint64) error {
	err := repository.GetDB().Session(&gorm.Session{AllowGlobalUpdate: true}).Table(model.GroupMessage{}.TableName()).
		Where("group_id = ?", groupID).Update("is_read", 1).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *MessageDal) GetLastMessageInUser(c *gin.Context, sendID, receiveID uint64) (string, error) {
	userMessage := &model.UserMessage{}
	err := repository.GetDB().Table(userMessage.TableName()).Where("send_user_id = ? and receive_user_id", sendID, receiveID).
		Order("created_at desc").First(userMessage).Error
	if err != nil {
		return "", err
	}
	return userMessage.Content, nil
}

func (ins *MessageDal) GetLastMessageInGroup(c *gin.Context, groupID uint64) (string, error) {
	groupMessage := &model.GroupMessage{}
	err := repository.GetDB().Table(groupMessage.TableName()).Where("group_id = ?", groupID).
		Order("created_at desc").First(groupMessage).Error
	if err != nil {
		return "", err
	}
	return groupMessage.Content, nil
}
