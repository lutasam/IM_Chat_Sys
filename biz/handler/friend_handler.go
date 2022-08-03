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
		r.GET("/find_groups/:input_str", friendController.FindGroups)
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

}

func (ins *FriendController) AddFriend(c *gin.Context) {

}

func (ins *FriendController) AddFriendsInGroup(c *gin.Context) {

}

func (ins *FriendController) FindFriends(c *gin.Context) {

}

func (ins *FriendController) FindGroups(c *gin.Context) {

}
