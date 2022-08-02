package service

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/bo"
	"github.com/lutasam/chat/biz/common"
	"github.com/lutasam/chat/biz/dal"
	"github.com/lutasam/chat/biz/model"
	"github.com/lutasam/chat/biz/utils"
	"github.com/lutasam/chat/biz/vo"
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
	_, err := dal.GetUserDal().GetUserByID(c, sendID)
	if err != nil {
		return nil, err
	}
	_, err = dal.GetUserDal().GetUserByID(c, receiveID)
	if err != nil {
		return nil, err
	}
	messages, err := dal.GetMessageDal().GetUserMessages(c, sendID, receiveID)
	if err != nil {
		return nil, err
	}
	err = dal.GetMessageDal().UpdateUserMessages(c, sendID, receiveID)
	if err != nil {
		return nil, err
	}
	messageVOs, err := convert2UserMessagesVO(messages)
	if err != nil {
		return nil, err
	}
	return &bo.GetUserMessagesResponse{
		Messages: messageVOs,
	}, nil
}

func (ins *MessageService) GetGroupMessages(c *gin.Context, groupID uint64) (*bo.GetGroupMessagesResponse, error) {
	group, err := dal.GetGroupDal().GetGroupByID(c, groupID)
	if err != nil {
		return nil, err
	}
	if group.ID == 0 {
		return nil, common.GROUPNOTEXIST
	}
	messages, err := dal.GetMessageDal().GetGroupMessages(c, groupID)
	if err != nil {
		return nil, err
	}
	err = dal.GetMessageDal().UpdateGroupMessages(c, groupID)
	if err != nil {
		return nil, err
	}
	messageVOs, err := convert2GroupMessagesVO(messages)
	if err != nil {
		return nil, err
	}
	return &bo.GetGroupMessagesResponse{
		Messages: messageVOs,
	}, nil
}

func (ins *MessageService) GetUserMessageNum(c *gin.Context, sendID, receiveID uint64) (*bo.GetUserMessageNumResponse, error) {
	_, err := dal.GetUserDal().GetUserByID(c, sendID)
	if err != nil {
		return nil, err
	}
	_, err = dal.GetUserDal().GetUserByID(c, receiveID)
	if err != nil {
		return nil, err
	}
	cnt, err := dal.GetMessageDal().GetUserMessageNum(c, sendID, receiveID)
	if err != nil {
		return nil, err
	}
	if cnt > common.MAXMESSAGENUM {
		cnt = common.MAXMESSAGENUM
	}
	return &bo.GetUserMessageNumResponse{
		Count: int(cnt),
	}, nil
}

func (ins *MessageService) GetGroupMessageNum(c *gin.Context, groupID uint64) (*bo.GetGroupMessageNumResponse, error) {
	_, err := dal.GetGroupDal().GetGroupByID(c, groupID)
	if err != nil {
		return nil, err
	}
	cnt, err := dal.GetMessageDal().GetGroupMessageNum(c, groupID)
	if err != nil {
		return nil, err
	}
	if cnt > common.MAXMESSAGENUM {
		cnt = common.MAXMESSAGENUM
	}
	return &bo.GetGroupMessageNumResponse{
		Count: int(cnt),
	}, nil
}

func (ins *MessageService) SendUserMessage(c *gin.Context, req *bo.SendUserMessageRequest, sendID uint64) error {
	receiveID, err := utils.ParseString2Uint64(req.ReceiveUserID)
	if err != nil {
		return err
	}
	user, err := dal.GetUserDal().GetUserByID(c, receiveID)
	if err != nil {
		return err
	}
	if user.ID == 0 {
		return common.USERDOESNOTEXIST
	}
	user, err = dal.GetUserDal().GetUserByID(c, sendID)
	if err != nil {
		return err
	}
	var contentEncrypt string
	contentEncrypt, err = utils.AesEncrypt(req.Content)
	err = dal.GetMessageDal().CreateUserMessage(c, &model.UserMessage{
		ID:            utils.GenerateMessageID(),
		SendUserID:    sendID,
		ReceiveUserID: receiveID,
		Content:       contentEncrypt,
		Name:          user.NickName,
		Avatar:        user.Avatar,
		IsRead:        false,
	})
	if err != nil {
		return err
	}
	return nil
}

func (ins *MessageService) SendGroupMessage(c *gin.Context, req *bo.SendGroupMessageRequest, sendID uint64) error {
	groupID, err := utils.ParseString2Uint64(req.GroupID)
	if err != nil {
		return err
	}
	_, err = dal.GetGroupDal().GetGroupByID(c, groupID)
	if err != nil {
		return err
	}
	user, err := dal.GetUserDal().GetUserByID(c, sendID)
	if err != nil {
		return err
	}
	contentEncrypt, err := utils.AesEncrypt(req.Content)
	err = dal.GetMessageDal().CreateGroupMessage(c, &model.GroupMessage{
		ID:      utils.GenerateMessageID(),
		GroupID: groupID,
		UserID:  sendID,
		Content: contentEncrypt,
		Name:    user.NickName,
		Avatar:  user.Avatar,
		IsRead:  false,
	})
	if err != nil {
		return err
	}
	return nil
}

func convert2UserMessagesVO(messages []*model.UserMessage) ([]*vo.UserMessagesVO, error) {
	var messageVOs []*vo.UserMessagesVO
	for _, message := range messages {
		contentDecrypt, err := utils.AesDecrypt(message.Content)
		if err != nil {
			return nil, err
		}
		messageVOs = append(messageVOs, &vo.UserMessagesVO{
			Name:    message.Name,
			Avatar:  message.Avatar,
			Content: contentDecrypt,
		})
	}
	return messageVOs, nil
}

func convert2GroupMessagesVO(messages []*model.GroupMessage) ([]*vo.GroupMessagesVO, error) {
	var messageVOs []*vo.GroupMessagesVO
	for _, message := range messages {
		contentDecrypt, err := utils.AesDecrypt(message.Content)
		if err != nil {
			return nil, err
		}
		messageVOs = append(messageVOs, &vo.GroupMessagesVO{
			Name:    message.Name,
			Avatar:  message.Avatar,
			Content: contentDecrypt,
		})
	}
	return messageVOs, nil
}
