package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/bo"
	"github.com/lutasam/chat/biz/common"
	"github.com/lutasam/chat/biz/dal"
	"github.com/lutasam/chat/biz/model"
	"github.com/lutasam/chat/biz/utils"
	"sync"
)

type UserService struct{}

var (
	userService     *UserService
	userServiceOnce sync.Once
)

func GetUserService() *UserService {
	userServiceOnce.Do(func() {
		userService = &UserService{}
	})
	return userService
}

func (ins *UserService) GetUserDetail(c *gin.Context, userID uint64) (*bo.GetUserDetailResponse, error) {
	user, err := dal.GetUserDal().GetUserByID(c, userID)
	if err != nil {
		return nil, err
	}
	return &bo.GetUserDetailResponse{
		ID:        utils.ParseUint642String(user.ID),
		Account:   user.Account,
		NickName:  user.NickName,
		Avatar:    user.Avatar,
		Sign:      user.Sign,
		Status:    common.ParseInt2Status(user.Status).String(),
		CreatedAt: utils.ParseTime2DateString(user.CreatedAt),
	}, nil
}

func (ins *UserService) UpdateUserInfo(c *gin.Context, req *bo.UpdateUserInfoRequest, userID uint64) error {
	if req.Avatar != "" && !utils.IsValidURL(req.Avatar) {
		return common.USERINPUTERROR
	}
	if req.NickName == "" {
		req.NickName = common.DEFAULTNICKNAME
	}
	if req.Avatar == "" {
		req.Avatar = common.DEFAULTAVATARURL
	}
	if req.Sign == "" {
		req.Sign = common.DEFAULTSIGN
	}
	_, err := dal.GetUserDal().GetUserByID(c, userID)
	if err != nil {
		return err
	}
	code, err := utils.EncryptPassword(req.Password)
	if err != nil {
		return err
	}
	err = dal.GetUserDal().UpdateUser(c, &model.User{
		ID:       userID,
		Password: code,
		NickName: req.NickName,
		Avatar:   req.Avatar,
		Sign:     req.Sign,
		Status:   common.ParseString2Status(req.Status).Int(),
	})
	if err != nil {
		return err
	}
	return nil
}
