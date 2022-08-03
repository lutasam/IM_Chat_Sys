package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/bo"
	"github.com/lutasam/chat/biz/common"
	"github.com/lutasam/chat/biz/dal"
	"github.com/lutasam/chat/biz/model"
	"github.com/lutasam/chat/biz/utils"
	"github.com/lutasam/chat/biz/vo"
	"sync"
)

type FriendService struct{}

var (
	friendService     *FriendService
	friendServiceOnce sync.Once
)

func GetFriendService() *FriendService {
	friendServiceOnce.Do(func() {
		friendService = &FriendService{}
	})
	return friendService
}

func (ins *FriendService) GetAllFriends(c *gin.Context, userID uint64) (*bo.GetAllFriendsResponse, error) {
	_, err := dal.GetUserDal().GetUserByID(c, userID)
	if err != nil {
		return nil, err
	}
	friends, err := dal.GetUserDal().GetUserFriends(c, userID)
	if err != nil {
		return nil, err
	}
	friendVos, err := convertFriends2FriendVOs(c, friends, userID)
	if err != nil {
		return nil, err
	}
	return &bo.GetAllFriendsResponse{
		Total:   len(friendVos),
		Friends: friendVos,
	}, nil
}

func (ins *FriendService) GetFriendDetail(c *gin.Context, friendID uint64) (*bo.GetFriendDetailResponse, error) {
	friend, err := dal.GetUserDal().GetUserByID(c, friendID)
	if err != nil {
		return nil, err
	}
	return &bo.GetFriendDetailResponse{
		ID:        utils.ParseUint642String(friend.ID),
		Account:   friend.Account,
		NickName:  friend.NickName,
		Avatar:    friend.Avatar,
		Sign:      friend.Sign,
		Status:    friend.Status,
		CreatedAt: utils.ParseTime2DateString(friend.CreatedAt),
	}, nil
}

func (ins *FriendService) AddFriend(c *gin.Context, userID, friendID uint64) error {
	_, err := dal.GetUserDal().GetUserByID(c, userID)
	if err != nil {
		return err
	}
	_, err = dal.GetUserDal().GetUserByID(c, friendID)
	if err != nil {
		return err
	}
	isFriend, err := isUserFriend(c, userID, friendID)
	if err != nil {
		return err
	}
	if isFriend {
		return common.HAVEBEENFRIEND
	}
	err = dal.GetUserDal().AddFriend(c, userID, friendID)
	if err != nil {
		return err
	}
	return nil
}

func (ins *FriendService) AddFriendInGroup(c *gin.Context, userID, groupID uint64) error {

}

func (ins *FriendService) FindFriends(c *gin.Context, inputStr string) error {

}

func (ins *FriendService) FindGroups(c *gin.Context, inputStr string) error {

}

func convertFriends2FriendVOs(c *gin.Context, friends []*model.User, userID uint64) ([]*vo.FriendVO, error) {
	var friendVos []*vo.FriendVO
	for _, friend := range friends {
		messageNum, err := dal.GetMessageDal().GetUserMessageNum(c, userID, friend.ID)
		if err != nil {
			return nil, err
		}
		lastMessage, err := dal.GetMessageDal().GetLastMessageInUser(c, userID, friend.ID)
		if err != nil {
			return nil, err
		}
		friendVos = append(friendVos, &vo.FriendVO{
			Nickname:    friend.NickName,
			Avatar:      friend.Avatar,
			MessageNum:  int(messageNum),
			LastMessage: lastMessage,
		})
	}
	return friendVos, nil
}

func isUserFriend(c *gin.Context, userID, friendID uint64) (bool, error) {
	friend, err := dal.GetUserDal().GetUserFriendByID(c, userID, friendID)
	if err != nil {
		return false, err
	}
	return friend.ID != 0, nil
}
