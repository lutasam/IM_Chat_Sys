package service

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetMessageService(t *testing.T) {
	Convey("TestGetMessageService", t, func() {
		service1 := GetMessageService()
		service2 := GetMessageService()
		So(service1, ShouldEqual, service2)
	})
}

func TestMessageService_GetGroupMessageNum(t *testing.T) {

}

func TestMessageService_GetGroupMessages(t *testing.T) {

}

func TestMessageService_GetUserMessageNum(t *testing.T) {

}

func TestMessageService_GetUserMessages(t *testing.T) {

}

func TestMessageService_SendGroupMessage(t *testing.T) {

}

func TestMessageService_SendUserMessage(t *testing.T) {

}
