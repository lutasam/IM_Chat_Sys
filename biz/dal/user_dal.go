package dal

import (
	"encoding/json"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/common"
	"github.com/lutasam/chat/biz/model"
	"github.com/lutasam/chat/biz/repository"
	"github.com/lutasam/chat/biz/utils"
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
	user := &model.User{}
	// try to get user detail from cache
	key := common.USERDETAILPREFIX + utils.ParseUint642String(userID)
	resp, err := repository.GetRedisDB().Get(key).Result()
	if err == nil {
		err := json.Unmarshal([]byte(resp), user)
		if err != nil {
			return nil, err
		}
		return user, nil
	}

	err = repository.GetDB().Table(user.TableName()).Where("id = ?", userID).Find(user).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if user.ID == 0 {
		return nil, common.USERDOESNOTEXIST
	}

	// send user detail to redis
	go func() {
		val, _ := json.Marshal(user)
		_, err := repository.GetRedisDB().Set(key, string(val), time.Hour*24).Result()
		if err != nil {
			// should log redis error, but no need to stop the program
		}
	}()

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

	// send user detail to redis
	go func() {
		val, _ := json.Marshal(user)
		_, err := repository.GetRedisDB().Set(common.USERDETAILPREFIX+utils.ParseUint642String(user.ID), string(val), time.Hour*24).Result()
		if err != nil {
			// should log redis error, but no need to stop the program
		}
	}()

	return user, nil
}

func (ins *UserDal) GetUsersByID(c *gin.Context, userID uint64) ([]*model.User, error) {
	var users []*model.User
	err := repository.GetDB().Table(model.User{}.TableName()).Where("id like ?", strconv.FormatUint(userID, 10)+"%").Find(&users).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return users, nil
}

func (ins *UserDal) GetUsersByAccount(c *gin.Context, account string) ([]*model.User, error) {
	var users []*model.User
	err := repository.GetDB().Table(model.User{}.TableName()).Where("account like ?", account+"%").Find(&users).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return users, nil
}

func (ins *UserDal) GetUsersByNickname(c *gin.Context, nickname string) ([]*model.User, error) {
	var users []*model.User
	err := repository.GetDB().Table(model.User{}.TableName()).Where("nickname like ?", nickname+"%").Find(&users).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return users, nil
}

// GetUserByAccountWithoutExistCheck this func will not check whether the user exist.
// in other word, if user does not exist, it will return an empty user.
func (ins *UserDal) GetUserByAccountWithoutExistCheck(c *gin.Context, account string) (*model.User, error) {
	user := &model.User{}
	err := repository.GetDB().Table(user.TableName()).Where("account = ?", account).Find(user).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return user, nil
}

func (ins *UserDal) UpdateUser(c *gin.Context, user *model.User) error {
	err := repository.GetDB().Table(user.TableName()).Updates(user).Error
	if err != nil {
		return common.DATABASEERROR
	}

	// delete cache in redis
	go func() {
		_, err = repository.GetRedisDB().Del(common.USERDETAILPREFIX + utils.ParseUint642String(user.ID)).Result()
		if err != nil {
			//
		}
	}()
	return nil
}

func (ins *UserDal) UpdateUserLoginInfo(c *gin.Context, userID uint64, ip string, port int) error {
	user := &model.User{
		IP:   ip,
		Port: port,
	}
	err := repository.GetDB().Table(model.User{}.TableName()).Where("id = ?", userID).Updates(user).Error
	if err != nil {
		return common.DATABASEERROR
	}

	// delete cache in redis
	go func() {
		_, err = repository.GetRedisDB().Del(common.USERDETAILPREFIX + utils.ParseUint642String(user.ID)).Result()
		if err != nil {
			//
		}
	}()
	return nil
}

func (ins *UserDal) GetUsersByIDs(c *gin.Context, ids []uint64) ([]*model.User, error) {
	var users []*model.User
	err := repository.GetDB().Table(model.User{}.TableName()).Where("id in ?", ids).Find(&users).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if len(users) != len(ids) {
		return nil, common.USERDOESNOTEXIST
	}
	return users, nil
}

func (ins *UserDal) GetUserFriends(c *gin.Context, userID uint64) ([]*model.User, error) {
	var users []*model.User
	err := repository.GetDB().Raw("select a.* from users as a, users_friends as b "+
		"where b.user_id = ? and b.friend_id = a.id", userID).Scan(&users).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return users, nil
}

func (ins *UserDal) AddFriend(c *gin.Context, userID, friendID uint64) error {
	err := repository.GetDB().Exec("insert into users_friends values(?, ?)", userID, friendID).Error
	if err != nil {
		return common.DATABASEERROR
	}
	err = repository.GetDB().Exec("insert into users_friends values(?, ?)", friendID, userID).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *UserDal) DeleteFriend(c *gin.Context, userID, friendID uint64) error {
	err := repository.GetDB().Exec("delete from users_friends "+
		"where user_id = ? and friend_id = ?", userID, friendID).Error
	if err != nil {
		return common.DATABASEERROR
	}
	err = repository.GetDB().Exec("delete from users_friends "+
		"where user_id = ? and friend_id = ?", friendID, userID).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *UserDal) GetUserFriendByID(c *gin.Context, userID, friendID uint64) (*model.User, error) {
	friend := &model.User{}
	err := repository.GetDB().
		Raw("select c.* from users as a, users_friends as b, users as c "+
			"where a.id = b.user_id and b.friend_id = c.id and a.id = ? and c.id = ?", userID, friendID).
		Scan(friend).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return friend, nil
}

func (ins *UserDal) AddFriendInGroup(c *gin.Context, userID, groupID uint64) error {
	groupMembers, err := GetGroupDal().GetGroupMembers(c, groupID)
	if err != nil {
		return err
	}
	err = repository.GetDB().Table(model.User{}.TableName()).Where("id = ?", userID).Association("Friends").
		Append(groupMembers)
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}
