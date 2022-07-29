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
