package service

import (
	"github.com/lutasam/chat/biz/dal"
	"github.com/lutasam/chat/biz/model"
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
	if req.Account == "" || req.Password == "" || !utils.IsValidIP(req.IP) ||
		!utils.IsValidPort(req.Port) {
		return nil, common.USERINPUTERROR
	}
	user, err := dal.GetUserDal().GetUserByAccount(c, req.Account)
	if err != nil {
		return nil, err
	}
	err = utils.ValidatePassword(user.Password, req.Password)
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

func (ins *LoginService) DoRegister(c *gin.Context, req *bo.RegisterRequest) (*bo.RegisterResponse, error) {
	if req.Account == "" || req.Password == "" || req.IP == "" ||
		!utils.IsValidPort(req.Port) || req.NickName == "" ||
		!utils.IsValidIP(req.IP) || !utils.IsValidURL(req.Avatar) {
		return nil, common.USERINPUTERROR
	}
	user, err := dal.GetUserDal().GetUserByAccount(c, req.Account)
	if err != nil {
		return nil, err
	}
	if user.ID != 0 {
		return nil, common.USEREXISTED
	}

	user = generateNewUser(req)
	var token string
	err = dal.GetUserDal().CreateUser(c, user)
	if err != nil {
		return nil, err
	}
	token, err = utils.GenerateJWTInUser(user)
	if err != nil {
		return nil, err
	}
	return &bo.RegisterResponse{
		Account: user.TableName(),
		Token:   token,
	}, nil
}

func generateNewUser(req *bo.RegisterRequest) *model.User {
	return &model.User{
		ID:       utils.GetID(),
		Account:  req.Account,
		Password: req.Password,
		NickName: req.NickName,
		Avatar:   req.Avatar,
		Sign:     req.Sign,
		IP:       req.IP,
		Port:     req.Port,
		Status:   common.ONLINE.Int(),
	}
}
