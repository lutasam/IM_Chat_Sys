package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/common"
	"github.com/lutasam/chat/biz/model"
	"github.com/lutasam/chat/biz/repository"
	"github.com/lutasam/chat/biz/utils"
	"sync"
)

type UserDal struct{}

var (
	userDal     *UserDal
	userDalOnce sync.Once
)

func GetUserDal() *UserDal {
	userDalOnce.Do(func() {
		userDal = &UserDal{}
	})
	return userDal
}

func (ins *UserDal) CreateUser(c *gin.Context, user *model.User) error {
	var err error
	user.Password, err = utils.EncryptPassword(user.Password)
	if err != nil {
		return err
	}
	err = repository.GetDB().Table(user.TableName()).Create(user).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *UserDal) GetUserByID(c *gin.Context, userID uint64) (*model.User, error) {
	var user *model.User
	err := repository.GetDB().Table(user.TableName()).Where("id = ?", userID).Find(user).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if user == nil {
		return nil, common.USERDOESNOTEXIST
	}
	return user, nil
}

func (ins *UserDal) GetUserByAccount(c *gin.Context, account string) (*model.User, error) {
	user := &model.User{}
	err := repository.GetDB().Table(user.TableName()).Where("account = ?", account).Find(user).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if user.ID == 0 {
		return nil, common.USERDOESNOTEXIST
	}
	return user, nil
}

func (ins *UserDal) UpdateUserByID(c *gin.Context, user *model.User) error {
	err := repository.GetDB().Table(user.TableName()).Updates(user).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}
