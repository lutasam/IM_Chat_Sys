package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/bo"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetUserService(t *testing.T) {
	Convey("TestGetUserService", t, func() {
		userService1 := GetUserService()
		userService2 := GetUserService()
		So(userService1, ShouldNotBeNil)
		So(userService1, ShouldEqual, userService2)
	})
}

func TestUserService_GetUserDetail(t *testing.T) {
	Convey("TestUserService_GetUserDetail", t, func() {
		Convey("normal", func() {
			c := &gin.Context{}
			resp, err := GetUserService().GetUserDetail(c, 0)
			So(err, ShouldBeNil)
			So(resp.ID, ShouldEqual, 0)
		})
	})
}

func TestUserService_UpdateUserInfo(t *testing.T) {
	Convey("TestUserService_UpdateUserInfo", t, func() {
		c := &gin.Context{}
		Convey("normal", func() {
			err := GetUserService().UpdateUserInfo(c, &bo.UpdateUserInfoRequest{
				Password: "12345",
				NickName: "test",
				Avatar:   "http://baidu.com/test.png",
				Sign:     "test",
				Status:   "ONLINE",
			}, 0)
			So(err, ShouldBeNil)
		})
		Convey("not fill all change options", func() {
			err := GetUserService().UpdateUserInfo(c, &bo.UpdateUserInfoRequest{
				Password: "12345",
				Avatar:   "http://baidu.com/test.png",
				Sign:     "test",
			}, 0)
			So(err, ShouldBeNil)
		})
		Convey("input format error", func() {
			err := GetUserService().UpdateUserInfo(c, &bo.UpdateUserInfoRequest{
				Avatar: "test.png",
				Status: "NOTONLINE",
			}, 0)
			So(err, ShouldNotBeNil)
		})
	})
}
