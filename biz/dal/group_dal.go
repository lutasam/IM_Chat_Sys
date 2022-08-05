package dal

import (
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/common"
	"github.com/lutasam/chat/biz/model"
	"github.com/lutasam/chat/biz/repository"
)

type GroupDal struct{}

var (
	groupDal     *GroupDal
	groupDalOnce sync.Once
)

func GetGroupDal() *GroupDal {
	groupDalOnce.Do(func() {
		groupDal = &GroupDal{}
	})
	return groupDal
}

func (ins *GroupDal) GetGroupByID(c *gin.Context, groupID uint64) (*model.Group, error) {
	group := &model.Group{}
	err := repository.GetDB().Table(group.TableName()).Preload("Users").
		Where("id = ?", groupID).Find(&group).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if group.ID == 0 {
		return nil, common.GROUPNOTEXIST
	}
	return group, nil
}

func (ins *GroupDal) GetGroupByName(c *gin.Context, name string) (*model.Group, error) {
	group := &model.Group{}
	err := repository.GetDB().Table(group.TableName()).Preload("Users").
		Where("name = ?", name).Find(&group).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if group.ID == 0 {
		return nil, common.GROUPNOTEXIST
	}
	return group, err
}

func (ins *GroupDal) GetGroupsByID(c *gin.Context, groupID uint64) ([]*model.Group, error) {
	var groups []*model.Group
	err := repository.GetDB().Table(model.Group{}.TableName()).
		Where("id like ?", strconv.FormatUint(groupID, 10)+"%").Find(&groups).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return groups, nil
}

func (ins *GroupDal) GetGroupsByName(c *gin.Context, name string) ([]*model.Group, error) {
	var groups []*model.Group
	err := repository.GetDB().Table(model.Group{}.TableName()).
		Where("name like ?", name+"%").Find(&groups).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return groups, nil
}

func (ins *GroupDal) GetUserGroups(c *gin.Context, userID uint64) ([]*model.Group, error) {
	var groups []*model.Group
	err := repository.GetDB().Raw("select b.* from users_groups as a, `groups` as b "+
		"where a.user_id = ? and a.group_id = b.id", userID).Scan(&groups).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return groups, nil
}

func (ins *GroupDal) CreateGroup(c *gin.Context, group *model.Group) error {
	err := repository.GetDB().Table(group.TableName()).Create(group).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *GroupDal) UpdateGroup(c *gin.Context, group *model.Group) error {
	err := repository.GetDB().Table(group.TableName()).Updates(group).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}

func (ins *GroupDal) GetGroupMembers(c *gin.Context, groupID uint64) ([]*model.User, error) {
	var groupMembers []*model.User
	err := repository.GetDB().Table(model.Group{}.TableName()).Where("id = ?", groupID).
		Association("Users").Find(&groupMembers).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return groupMembers, nil
}

func (ins *GroupDal) AttendInGroup(c *gin.Context, groupID, userID uint64) error {
	err := repository.GetDB().Exec("insert into users_groups values(?, ?)", groupID, userID).Error
	if err != nil {
		return common.DATABASEERROR
	}
	return nil
}
