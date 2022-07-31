package bo

import "github.com/lutasam/chat/biz/vo"

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
	ID       string `json:"id"`
	Account  string `json:"account"`
	NickName string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Sign     string `json:"sign"`
	Status   int    `json:"status"`
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
