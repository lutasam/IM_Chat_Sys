package handler

import (
	"github.com/gin-gonic/gin"
)

type GroupController struct{}

func RegisterGroupRouter(r *gin.RouterGroup) {
	groupController := &GroupController{}
	{
		r.POST("/create_group", groupController.CreateGroup)
		r.POST("/update_group", groupController.UpdateGroup)
		r.GET("/get_group_detail/:group_id", groupController.GetGroupDetail)
		r.GET("/get_all_groups", groupController.GetAllGroups)
	}
}

func (ins *GroupController) CreateGroup(c *gin.Context) {

}

func (ins *GroupController) UpdateGroup(c *gin.Context) {

}

func (ins *GroupController) GetGroupDetail(c *gin.Context) {

}

func (ins *GroupController) GetAllGroups(c *gin.Context) {

}
