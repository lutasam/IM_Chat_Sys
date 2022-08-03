package bo

import (
	"github.com/lutasam/chat/biz/vo"
)

type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type LoginResponse struct {
	Account string `json:"account"`
	Token   string `json:"token"`
}

type RegisterResponse struct {
	Account string `json:"account"`
	Token   string `json:"token"`
}

type GetUserDetailResponse struct {
	ID        string `json:"id"`
	Account   string `json:"account"`
	NickName  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	Sign      string `json:"sign"`
	Status    int    `json:"status"`
	CreatedAt string `json:"created_at"`
}

type GetUserMessagesResponse struct {
	Messages []*vo.UserMessagesVO `json:"messages"`
}

type GetGroupMessagesResponse struct {
	Messages []*vo.GroupMessagesVO `json:"messages"`
}

type GetUserMessageNumResponse struct {
	Count int `json:"count"`
}

type GetGroupMessageNumResponse struct {
	Count int `json:"count"`
}

type GetGroupDetailResponse struct {
	ID        uint64      `json:"id"`
	Name      string      `json:"name"`
	Describe  string      `json:"describe"`
	Avatar    string      `json:"avatar"`
	MemberNum int         `json:"member_num"`
	AdminUser *vo.UserVO  `json:"admin_id"`
	Tags      []*vo.TagVO `json:"tages"`
	CreatedAt string      `json:"created_at"`
}

type GetAllGroupsResponse struct {
	Total  int                      `json:"total"`
	Groups []*vo.GroupWithMessageVO `json:"groups"`
}

type GetAllFriendsResponse struct {
	Total   int            `json:"total"`
	Friends []*vo.FriendVO `json:"friends"`
}

type GetFriendDetailResponse struct {
	ID        string `json:"id"`
	Account   string `json:"account"`
	NickName  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	Sign      string `json:"sign"`
	Status    int    `json:"status"`
	CreatedAt string `json:"created_at"`
}
