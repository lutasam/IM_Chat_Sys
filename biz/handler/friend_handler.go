package handler

import "github.com/gin-gonic/gin"

type FriendController struct{}

func RegisterFriendRouter(r *gin.RouterGroup) {
	friendController := &FriendController{}
	{
		r.GET("/get_all_friends", friendController.GetAllFriends)
		r.GET("/get_friend_detail/:friend_id", friendController.GetFriendDetail)
		r.GET("/add_friend/:friend_id", friendController.AddFriend)
		r.GET("/add_friends_in_group/:group_id", friendController.AddFriendsInGroup)
		r.GET("/find_friend_by_nickname/:nickname", friendController.FindFriend)
		r.GET("/find_group_by_name/:name", friendController.FindGroup)
	}
}

func (ins *FriendController) GetAllFriends(c *gin.Context) {

}

func (ins *FriendController) GetFriendDetail(c *gin.Context) {

}

func (ins *FriendController) AddFriend(c *gin.Context) {

}

func (ins *FriendController) AddFriendsInGroup(c *gin.Context) {

}

func (ins *FriendController) FindFriend(c *gin.Context) {

}

func (ins *FriendController) FindGroup(c *gin.Context) {

}
