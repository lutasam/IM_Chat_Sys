package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/common"
	"github.com/lutasam/chat/biz/service"
	"github.com/lutasam/chat/biz/utils"
)

type FriendController struct{}

func RegisterFriendRouter(r *gin.RouterGroup) {
	friendController := &FriendController{}
	{
		r.GET("/get_all_friends", friendController.GetAllFriends)
		r.GET("/get_friend_detail/:friend_id", friendController.GetFriendDetail)
		r.GET("/add_friend/:friend_id", friendController.AddFriend)
		r.GET("/add_friends_in_group/:group_id", friendController.AddFriendsInGroup)
		r.GET("/find_friends/:input_str", friendController.FindFriends)
	}
}

func (ins *FriendController) GetAllFriends(c *gin.Context) {
	jwtStruct, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		utils.ResponseClientError(c, common.USERNOTLOGIN)
		return
	}
	resp, err := service.GetFriendService().GetAllFriends(c, jwtStruct.UserID)
	if err != nil {
		if errors.Is(err, common.USERDOESNOTEXIST) {
			utils.ResponseClientError(c, common.USERDOESNOTEXIST)
			return
		} else if errors.Is(err, common.DATANOTFOUND) {
			utils.ResponseClientError(c, common.DATANOTFOUND)
			return
		} else {
			utils.ResponseServerError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *FriendController) GetFriendDetail(c *gin.Context) {
	idstr := c.Param("friend_id")
	friendID, err := utils.ParseString2Uint64(idstr)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	resp, err := service.GetFriendService().GetFriendDetail(c, friendID)
	if err != nil {
		if errors.Is(err, common.USERDOESNOTEXIST) {
			utils.ResponseClientError(c, common.USERDOESNOTEXIST)
			return
		} else {
			utils.ResponseServerError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}

func (ins *FriendController) AddFriend(c *gin.Context) {
	jwtStruct, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		utils.ResponseClientError(c, common.USERNOTLOGIN)
		return
	}
	idstr := c.Param("friend_id")
	friendID, err := utils.ParseString2Uint64(idstr)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	err = service.GetFriendService().AddFriend(c, jwtStruct.UserID, friendID)
	if err != nil {
		if errors.Is(err, common.USERDOESNOTEXIST) {
			utils.ResponseClientError(c, common.USERDOESNOTEXIST)
			return
		} else {
			utils.ResponseServerError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, nil)
}

func (ins *FriendController) AddFriendsInGroup(c *gin.Context) {
	jwtStruct, err := utils.GetCtxUserInfoJWT(c)
	if err != nil {
		utils.ResponseClientError(c, common.USERNOTLOGIN)
		return
	}
	idstr := c.Param("group_id")
	groupID, err := utils.ParseString2Uint64(idstr)
	if err != nil {
		utils.ResponseClientError(c, common.USERINPUTERROR)
		return
	}
	err = service.GetFriendService().AddFriendInGroup(c, jwtStruct.UserID, groupID)
	if err != nil {
		if errors.Is(err, common.USERDOESNOTEXIST) {
			utils.ResponseClientError(c, common.USERDOESNOTEXIST)
			return
		} else if errors.Is(err, common.GROUPNOTEXIST) {
			utils.ResponseClientError(c, common.GROUPNOTEXIST)
			return
		} else {
			utils.ResponseServerError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, nil)
}

func (ins *FriendController) FindFriends(c *gin.Context) {
	str := c.Param("input_str")
	resp, err := service.GetFriendService().FindFriends(c, str)
	if err != nil {
		if errors.Is(err, common.USERINPUTERROR) {
			utils.ResponseClientError(c, common.USERINPUTERROR)
			return
		} else {
			utils.ResponseServerError(c, common.UNKNOWNERROR)
			return
		}
	}
	utils.ResponseSuccess(c, resp)
}
