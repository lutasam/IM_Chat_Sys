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
	if user.ID == 0 {
		return nil, common.USERDOESNOTEXIST
	}
	return &bo.GetUserDetailResponse{
		ID:       utils.ParseUint642String(user.ID),
		Account:  user.Account,
		NickName: user.NickName,
		Avatar:   user.Avatar,
		Sign:     user.Sign,
		Status:   user.Status,
	}, nil
}

func (ins *UserService) UpdateUserInfo(c *gin.Context, req *bo.UpdateUserInfoRequest) error {
	if req.Avatar != "" && !utils.IsValidURL(req.Avatar) {
		return common.USERINPUTERROR
	}
	jwtStruct, exist := c.Get("jwtStruct")
	if !exist {
		return common.USERNOTLOGIN
	}
	code, err := utils.EncryptPassword(req.Password)
	if err != nil {
		return err
	}
	err = dal.GetUserDal().UpdateUser(c, &model.User{
		ID:       jwtStruct.(*utils.JWTStruct).UserID,
		Password: code,
		NickName: req.NickName,
		Avatar:   req.Avatar,
		Sign:     req.Sign,
	})
	if err != nil {
		return err
	}
	return nil
}
