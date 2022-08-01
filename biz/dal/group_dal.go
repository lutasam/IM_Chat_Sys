package dal

import (
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
	err := repository.GetDB().Table(group.TableName()).Where("id = ?", groupID).Find(&group).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	return group, nil
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