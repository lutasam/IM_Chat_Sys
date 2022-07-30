package bo

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
	Account  string `json:"account"`
	NickName string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Sign     string `json:"sign"`
	Status   int    `json:"status"`
}
