package handler

import "github.com/gin-gonic/gin"

type GroupController struct{}

func RegisterGroupRouter(r *gin.RouterGroup) {
	groupController := &GroupController{}
	{
		r.POST("/create_group", groupController.CreateGroup)
		r.POST("/update_group", groupController.UpdateGroup)
		r.GET("/get_group/:group_id", groupController.GetGroup)

	}
}

func (ins *GroupController) CreateGroup(c *gin.Context) {

}

func (ins *GroupController) UpdateGroup(c *gin.Context) {

}

func (ins *GroupController) GetGroup(c *gin.Context) {

}
