package bo

type LoginRequest struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
	IP       string `json:"ip" binding:"required"`
	Port     int    `json:"port" binding:"required"`
}

type RegisterRequest struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
	NickName string `json:"nickname" binding:"required"`
	Avatar   string `json:"avatar" binding:"-"`
	Sign     string `json:"sign" binding:"-"`
	IP       string `json:"ip" binding:"required"`
	Port     int    `json:"port" binding:"required"`
}

type UpdateUserInfoRequest struct {
	Password string `json:"password" binding:"-"`
	NickName string `json:"nickname" binding:"-"`
	Avatar   string `json:"avatar" binding:"-"`
	Sign     string `json:"sign" binding:"-"`
}

type SendUserMessageRequest struct {
	ReceiveUserID string `json:"receive_user_id" binding:"required"`
	Content       string `json:"content" binging:"required"`
}

type SendGroupMessageRequest struct {
	GroupID string `json:"group_id" binding:"required"`
	Content string `json:"content" binging:"required"`
}
