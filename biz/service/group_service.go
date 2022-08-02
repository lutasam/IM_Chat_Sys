package service

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/bo"
	"github.com/lutasam/chat/biz/common"
	"github.com/lutasam/chat/biz/dal"
	"github.com/lutasam/chat/biz/model"
	"github.com/lutasam/chat/biz/utils"
	"github.com/lutasam/chat/biz/vo"
)

type GroupService struct{}

var (
	groupService     *GroupService
	groupServiceOnce sync.Once
)

func GetGroupService() *GroupService {
	groupServiceOnce.Do(func() {
		groupService = &GroupService{}
	})
	return groupService
}

func (ins *GroupService) CreateGroup(c *gin.Context, req *bo.CreateGroupRequest, creator_id uint64) error {
	if req.Avatar != "" && !utils.IsValidURL(req.Avatar) {
		return common.USERDOESNOTEXIST
	}
	group, err := convertCreateGroupRequest2Group(c, req, creator_id)
	if err != nil {
		return err
	}
	err = dal.GetGroupDal().CreateGroup(c, group)
	if err != nil {
		return err
	}
	return nil
}

func (ins *GroupService) UpdateGroup(c *gin.Context, req *bo.UpdateGroupRequest) error {
	if req.Avatar != "" && !utils.IsValidURL(req.Avatar) || req.Name == "" {
		return common.USERINPUTERROR
	}
	group, err := convertUpdateGroupRequest2Group(c, req)
	if err != nil {
		return err
	}
	err = dal.GetGroupDal().UpdateGroup(c, group)
	if err != nil {
		return err
	}
	return nil
}

func (ins *GroupService) GetGroupDetail(c *gin.Context, name string) (*bo.GetGroupDetailResponse, error) {
	group, err := dal.GetGroupDal().GetGroupByName(c, name)
	if err != nil {
		return nil, err
	}
	var adminUser *model.User
	adminUser, err = dal.GetUserDal().GetUserByID(c, group.AdminID)
	if err != nil {
		return nil, err
	}
	tagVOs := convertTags2TagVOs(group.Tages)
	return &bo.GetGroupDetailResponse{
		ID:        group.ID,
		Name:      group.Name,
		Describe:  group.Describe,
		Avatar:    group.Avatar,
		MemberNum: len(group.User),
		AdminUser: &vo.UserVO{
			Name:   adminUser.NickName,
			Avatar: adminUser.Avatar,
		},
		Tags:      tagVOs,
		CreatedAt: group.CreatedAt,
	}, nil
}

// TODO: not realize not use this func!
func (ins *GroupService) GetAllGroups(c *gin.Context, userID uint64) (*bo.GetAllGroupsResponse, error) {
	_, err := dal.GetUserDal().GetUserByID(c, userID)
	if err != nil {
		return nil, err
	}
	var groups []*model.Group
	groups, err = dal.GetGroupDal().GetUserGroups(c, userID)
	if err != nil {
		return nil, err
	}
	groupVOs := convertGroups2GroupWithMessageVOs(groups)
	return &bo.GetAllGroupsResponse{
		Total:  len(groupVOs),
		Groups: groupVOs,
	}, nil
}

// TODO: not realize not use this func!
func convertGroups2GroupWithMessageVOs(groups []*model.Group) []*vo.GroupWithMessageVO {
	var groupVOs []*vo.GroupWithMessageVO
	for _, group := range groups {
		groupVOs = append(groupVOs, &vo.GroupWithMessageVO{
			Name:        group.Name,
			Avatar:      group.Avatar,
			MessageNum:  0,
			LastMessage: "not realize",
		})
	}
}

func convertTags2TagVOs(tags []*model.Tag) []*vo.TagVO {
	var tagVOs []*vo.TagVO
	for _, tag := range tags {
		tagVOs = append(tagVOs, &vo.TagVO{
			Name: tag.Name,
		})
	}
	return tagVOs
}

func convertCreateGroupRequest2Group(c *gin.Context, req *bo.CreateGroupRequest, userID uint64) (*model.Group, error) {
	var ids []uint64
	for _, idstr := range req.UserIDs {
		id, err := utils.ParseString2Uint64(idstr)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	users, err := dal.GetUserDal().GetUsersByIDs(c, ids)
	if err != nil {
		return nil, err
	}
	var tags []*model.Tag
	tags, err = dal.GetTagDal().GetTagsByNames(c, req.TagNames)
	if err != nil {
		return nil, err
	}
	return &model.Group{
		ID:       utils.GenerateGroupID(),
		Name:     req.Name,
		Describe: req.Describe,
		Avatar:   req.Avatar,
		AdminID:  userID,
		User:     users,
		Tages:    tags,
	}, nil
}

func convertUpdateGroupRequest2Group(c *gin.Context, req *bo.UpdateGroupRequest) (*model.Group, error) {
	ids := make([]uint64, 0)
	for _, idstr := range req.UserIDs {
		id, err := utils.ParseString2Uint64(idstr)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	users, err := dal.GetUserDal().GetUsersByIDs(c, ids)
	if err != nil {
		return nil, err
	}
	var tags []*model.Tag
	tags, err = dal.GetTagDal().GetTagsByNames(c, req.TagNames)
	if err != nil {
		return nil, err
	}
	var id uint64
	id, err = utils.ParseString2Uint64(req.GroupID)
	if err != nil {
		return nil, err
	}
	return &model.Group{
		ID:       id,
		Name:     req.Name,
		Describe: req.Describe,
		Avatar:   req.Avatar,
		User:     users,
		Tages:    tags,
	}, nil
}
