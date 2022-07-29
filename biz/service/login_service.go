package service

import (
	"github.com/lutasam/chat/biz/dal"
	"github.com/lutasam/chat/biz/utils"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/bo"
	"github.com/lutasam/chat/biz/common"
)

type LoginService struct{}

var (
	loginService     *LoginService
	loginServiceOnce sync.Once
)

func GetLoginService() *LoginService {
	loginServiceOnce.Do(func() {
		loginService = &LoginService{}
	})
	return loginService
}

func (ins *LoginService) DoLogin(c *gin.Context, req *bo.LoginRequest) (*bo.LoginResponse, error) {
	if req.Account == "" || req.Password == "" {
		return nil, common.USERINPUTERROR
	}
	user, err := dal.GetUserDal().GetUserByAccount(c, req.Account)
	if err != nil {
		return nil, err
	}
	var token string
	token, err = utils.GenerateJWTInUser(user)
	if err != nil {
		return nil, err
	}
	return &bo.LoginResponse{
		Account: user.Account,
		Token:   token,
	}, nil
}
