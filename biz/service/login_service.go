package service

import (
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
	if req.Username == "" || req.Password == "" {
		return nil, common.USERINPUTERROR
	}

	return nil, nil
}
