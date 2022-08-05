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

func (ins *GroupService) CreateGroup(c *gin.Context, req *bo.CreateGroupRequest, userID uint64) error {
	if req.Avatar != "" && !utils.IsValidURL(req.Avatar) {
		return common.USERDOESNOTEXIST
	}
	if req.Avatar == "" {
		req.Avatar = common.DEFAULTAVATARURL
	}
	if req.Describe == "" {
		req.Describe = common.DEFAULTDESCRIPTION
	}
	user, err := dal.GetUserDal().GetUserByID(c, userID)
	if err != nil {
		return err
	}
	group, err := convertCreateGroupRequest2Group(c, req, userID)
	if err != nil {
		return err
	}
	group.Users = append(group.Users, user)
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
	if req.Avatar == "" {
		req.Avatar = common.DEFAULTAVATARURL
	}
	if req.Describe == "" {
		req.Describe = common.DEFAULTDESCRIPTION
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

func (ins *GroupService) GetGroupDetail(c *gin.Context, groupID uint64) (*bo.GetGroupDetailResponse, error) {
	group, err := dal.GetGroupDal().GetGroupByID(c, groupID)
	if err != nil {
		return nil, err
	}
	adminUser, err := dal.GetUserDal().GetUserByID(c, group.AdminID)
	if err != nil {
		return nil, err
	}
	tagVOs := convertTags2TagVOs(group.Tags)
	return &bo.GetGroupDetailResponse{
		ID:        group.ID,
		Name:      group.Name,
		Describe:  group.Describe,
		Avatar:    group.Avatar,
		MemberNum: len(group.Users),
		AdminUser: &vo.UserVO{
			Name:   adminUser.NickName,
			Avatar: adminUser.Avatar,
		},
		Tags:      tagVOs,
		CreatedAt: utils.ParseTime2DateString(group.CreatedAt),
	}, nil
}

// GetAllGroups
func (ins *GroupService) GetAllGroups(c *gin.Context, userID uint64) (*bo.GetAllGroupsResponse, error) {
	_, err := dal.GetUserDal().GetUserByID(c, userID)
	if err != nil {
		return nil, err
	}
	groups, err := dal.GetGroupDal().GetUserGroups(c, userID)
	if err != nil {
		return nil, err
	}
	groupVOs := convertGroups2GroupWithMessageVOs(c, groups)
	return &bo.GetAllGroupsResponse{
		Total:  len(groupVOs),
		Groups: groupVOs,
	}, nil
}

func (ins *GroupService) FindGroups(c *gin.Context, inputStr string) (*bo.FindGroupsResponse, error) {
	var groups []*model.Group
	maybeID, err := utils.ParseString2Uint64(inputStr)
	if err == nil {
		groupsByID, err := dal.GetGroupDal().GetGroupsByID(c, maybeID)
		if err != nil {
			return nil, err
		}
		groups = append(groups, groupsByID...)
	}
	groupsByName, err := dal.GetGroupDal().GetGroupsByName(c, inputStr)
	if err != nil {
		return nil, err
	}
	groups = append(groups, groupsByName...)
	groupVOs := convertGroups2GroupInSearchVOs(groups)
	if err != nil {
		return nil, err
	}
	return &bo.FindGroupsResponse{
		Total:  len(groupVOs),
		Groups: groupVOs,
	}, nil
}

func (ins *GroupService) AttendInGroup(c *gin.Context, groupID, userID uint64) error {
	_, err := dal.GetGroupDal().GetGroupByID(c, groupID)
	if err != nil {
		return err
	}
	_, err = dal.GetUserDal().GetUserByID(c, userID)
	if err != nil {
		return err
	}
	err = dal.GetGroupDal().AttendInGroup(c, groupID, userID)
	if err != nil {
		return err
	}
	return nil
}

func convertGroups2GroupInSearchVOs(groups []*model.Group) []*vo.GroupInSearchVO {
	groups = convert2NoReplicaGroupSlice(groups)
	var groupVos []*vo.GroupInSearchVO
	for _, group := range groups {
		groupVos = append(groupVos, &vo.GroupInSearchVO{
			ID:     utils.ParseUint642String(group.ID),
			Name:   group.Name,
			Avatar: group.Avatar,
		})
	}
	return groupVos
}

func convertGroups2GroupWithMessageVOs(c *gin.Context, groups []*model.Group) []*vo.GroupVO {
	var groupVOs []*vo.GroupVO
	for _, group := range groups {
		groupVOs = append(groupVOs, &vo.GroupVO{
			Name:   group.Name,
			Avatar: group.Avatar,
		})
	}
	return groupVOs
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
	tags, err := dal.GetTagDal().GetTagsByNames(c, req.TagNames)
	if err != nil {
		return nil, err
	}
	return &model.Group{
		ID:       utils.GenerateGroupID(),
		Name:     req.Name,
		Describe: req.Describe,
		Avatar:   req.Avatar,
		AdminID:  userID,
		Users:    users,
		Tags:     tags,
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
	tags, err := dal.GetTagDal().GetTagsByNames(c, req.TagNames)
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
		Users:    users,
		Tags:     tags,
	}, nil
}

func convert2NoReplicaGroupSlice(groups []*model.Group) []*model.Group {
	set := make(map[uint64]*model.Group)
	for _, group := range groups {
		if _, exist := set[group.ID]; !exist {
			set[group.ID] = group
		}
	}
	var groupNoReplica []*model.Group
	for _, group := range set {
		groupNoReplica = append(groupNoReplica, group)
	}
	return groupNoReplica
}
