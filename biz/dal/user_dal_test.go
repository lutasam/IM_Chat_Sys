package dal

import (
	"github.com/lutasam/chat/biz/model"
	"github.com/lutasam/chat/biz/utils"
	"testing"
)

func TestGetUserDal(t *testing.T) {

}

func TestUserDal_CreateUser(t *testing.T) {
	err := GetUserDal().CreateUser(nil, &model.User{
		ID:       utils.GetID(),
		Account:  "lutasam",
		Password: "123456",
		NickName: "lutasam",
		Avatar:   "test url",
		Sign:     "ni hao",
		IP:       "127.0.0.1",
		Port:     9000,
		Status:   0,
	})
	if err != nil {
		return
	}
}

func TestUserDal_GetUserByID(t *testing.T) {

}

func TestUserDal_GetUserByName(t *testing.T) {

}
