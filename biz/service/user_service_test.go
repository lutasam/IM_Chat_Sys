package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/bo"
	"reflect"
	"testing"
)

func TestGetUserService(t *testing.T) {
	tests := []struct {
		name string
		want *UserService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUserService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_GetUserDetail(t *testing.T) {
	type args struct {
		c      *gin.Context
		userID uint64
	}
	tests := []struct {
		name    string
		args    args
		want    *bo.GetUserDetailResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ins := &UserService{}
			got, err := ins.GetUserDetail(tt.args.c, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UpdateUserInfo(t *testing.T) {
	type args struct {
		c   *gin.Context
		req *bo.UpdateUserInfoRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ins := &UserService{}
			if err := ins.UpdateUserInfo(tt.args.c, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
