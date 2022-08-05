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
	friendVos := convertFriends2FriendVOs(c, friends, userID)
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

func (ins *FriendService) DeleteFriend(c *gin.Context, userID, friendID uint64) error {
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
	if !isFriend {
		return common.NOTFRIEND
	}
	err = dal.GetUserDal().DeleteFriend(c, userID, friendID)
	if err != nil {
		return err
	}
	return nil
}

func (ins *FriendService) AddFriendInGroup(c *gin.Context, userID, groupID uint64) error {
	_, err := dal.GetUserDal().GetUserByID(c, userID)
	if err != nil {
		return err
	}
	_, err = dal.GetGroupDal().GetGroupByID(c, groupID)
	if err != nil {
		return err
	}
	err = dal.GetUserDal().AddFriendInGroup(c, userID, groupID)
	if err != nil {
		return err
	}
	return nil
}

func (ins *FriendService) FindFriends(c *gin.Context, inputStr string) (*bo.FindFriendsResponse, error) {
	var friends []*model.User
	maybeID, err := utils.ParseString2Uint64(inputStr)
	if err == nil {
		friendsByID, err := dal.GetUserDal().GetUsersByID(c, maybeID)
		if err != nil {
			return nil, err
		}
		friends = append(friends, friendsByID...)
	}
	friendsByAccount, err := dal.GetUserDal().GetUsersByAccount(c, inputStr)
	if err != nil {
		return nil, err
	}
	friends = append(friends, friendsByAccount...)
	friendsByNickname, err := dal.GetUserDal().GetUsersByNickname(c, inputStr)
	if err != nil {
		return nil, err
	}
	friends = append(friends, friendsByNickname...)
	friendVOs := convertFriends2FriendInSearchVOs(c, friends)
	if err != nil {
		return nil, err
	}
	return &bo.FindFriendsResponse{
		Total:   len(friendVOs),
		Friends: friendVOs,
	}, nil
}

func convertFriends2FriendInSearchVOs(c *gin.Context, friends []*model.User) []*vo.FriendInSearchVO {
	friends = convert2NoReplicaFriendSlice(friends)
	var friendVos []*vo.FriendInSearchVO
	for _, friend := range friends {
		friendVos = append(friendVos, &vo.FriendInSearchVO{
			ID:       utils.ParseUint642String(friend.ID),
			Account:  friend.Account,
			Nickname: friend.NickName,
			Avatar:   friend.Avatar,
		})
	}
	return friendVos
}

func convert2NoReplicaFriendSlice(friends []*model.User) []*model.User {
	set := make(map[uint64]*model.User)
	for _, friend := range friends {
		if _, exist := set[friend.ID]; !exist {
			set[friend.ID] = friend
		}
	}
	var friendNoReplica []*model.User
	for _, friend := range set {
		friendNoReplica = append(friendNoReplica, friend)
	}
	return friendNoReplica
}

func convertFriends2FriendVOs(c *gin.Context, friends []*model.User, userID uint64) []*vo.FriendVO {
	var friendVos []*vo.FriendVO
	for _, friend := range friends {
		friendVos = append(friendVos, &vo.FriendVO{
			Nickname: friend.NickName,
			Avatar:   friend.Avatar,
		})
	}
	return friendVos
}

func isUserFriend(c *gin.Context, userID, friendID uint64) (bool, error) {
	friend, err := dal.GetUserDal().GetUserFriendByID(c, userID, friendID)
	if err != nil {
		return false, err
	}
	return friend.ID != 0, nil
}
