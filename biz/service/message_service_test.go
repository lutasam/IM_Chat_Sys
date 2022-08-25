package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/bo"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetMessageService(t *testing.T) {
	Convey("TestGetMessageService", t, func() {
		service1 := GetMessageService()
		service2 := GetMessageService()
		So(service1, ShouldNotBeNil)
		So(service1, ShouldEqual, service2)
	})
}

func TestMessageService_GetGroupMessageNum(t *testing.T) {
	Convey("TestMessageService_GetGroupMessageNum", t, func() {
		c := &gin.Context{}
		Convey("normal", func() {
			resp, err := GetMessageService().GetGroupMessageNum(c, 0)
			So(err, ShouldBeNil)
			So(resp, ShouldEqual, 0)
		})
		Convey("group_id not exist", func() {
			_, err := GetMessageService().GetGroupMessageNum(c, 0)
			So(err, ShouldNotBeNil)
		})
	})
}

func TestMessageService_GetGroupMessages(t *testing.T) {
	Convey("TestMessageService_GetGroupMessages", t, func() {
		c := &gin.Context{}
		Convey("normal", func() {
			resp, err := GetMessageService().GetGroupMessages(c, 0)
			So(err, ShouldBeNil)
			So(len(resp.Messages), ShouldEqual, 0)
		})
	})
}

func TestMessageService_GetUserMessageNum(t *testing.T) {
	Convey("TestMessageService_GetUserMessageNum", t, func() {
		c := &gin.Context{}
		Convey("normal", func() {
			resp, err := GetMessageService().GetUserMessageNum(c, 0, 1)
			So(err, ShouldBeNil)
			So(resp.Count, ShouldEqual, 0)
		})
	})
}

func TestMessageService_GetUserMessages(t *testing.T) {
	Convey("TestMessageService_GetUserMessages", t, func() {
		c := &gin.Context{}
		Convey("normal", func() {
			resp, err := GetMessageService().GetUserMessages(c, 0, 1)
			So(err, ShouldBeNil)
			So(len(resp.Messages), ShouldEqual, 0)
		})
	})
}

func TestMessageService_SendGroupMessage(t *testing.T) {
	Convey("TestMessageService_SendGroupMessage", t, func() {
		c := &gin.Context{}
		Convey("normal", func() {
			err := GetMessageService().SendGroupMessage(c, &bo.SendGroupMessageRequest{
				GroupID: "0",
				Content: "hello",
			}, 0)
			So(err, ShouldBeNil)
		})
	})
}

func TestMessageService_SendUserMessage(t *testing.T) {
	Convey("TestMessageService_SendUserMessage", t, func() {
		c := &gin.Context{}
		Convey("normal", func() {
			// err := GetMessageService().SendUserMessage(c, &bo.SendUserMessageRequest{
			// 	ReceiveUserID: "",
			// 	Content:       "",
			// })
		})
	})
}

func TestMessageService_GetAllMessages(t *testing.T) {
	Convey("TestMessageService_GetAllMessages", t, func() {
		c := &gin.Context{}
		Convey("normal", func() {

		})
	})
}
