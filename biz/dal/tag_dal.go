package dal

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/common"
	"github.com/lutasam/chat/biz/model"
	"github.com/lutasam/chat/biz/repository"
)

type TagDal struct{}

var (
	tagDal     *TagDal
	tagDalOnce sync.Once
)

func GetTagDal() *TagDal {
	tagDalOnce.Do(func() {
		tagDal = &TagDal{}
	})
	return tagDal
}

func (ins *TagDal) GetTagByName(c *gin.Context, name string) (*model.Tag, error) {
	tag := &model.Tag{}
	err := repository.GetDB().Table(tag.TableName()).Where("name = ?", name).Find(tag).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if tag.ID == 0 {
		return nil, common.DATANOTFOUND
	}
	return tag, nil
}

func (ins *TagDal) GetTagsByNames(c *gin.Context, names []string) ([]*model.Tag, error) {
	var tags []*model.Tag
	err := repository.GetDB().Table(model.Tag{}.TableName()).Where("name in ?", names).Find(&tags).Error
	if err != nil {
		return nil, common.DATABASEERROR
	}
	if len(tags) != len(names) {
		return nil, common.DATANOTFOUND
	}
	return tags, nil
}
