package common

const ISSUER = "LUTASAM"
const JWTSECRETSALT = "LUTASAM"
const PASSWORDSALT = "astaxie12798akljzmknm.ahkjkljl;k"
const EXPIRETIME = 86400000

const (
	STATUSOKCODE    = 200
	CLIENTERRORCODE = 404
	SERVERERRORCODE = 500
)

const (
	STATUSOKMSG    = "OK"
	CLIENTERRORMSG = "404 NOT FOUND"
	SERVERERRORMSG = "500 SERVER ERROR"
)

// Status 用户状态
type Status int

const (
	ONLINE     Status = iota + 1 // 在线
	BUSY                         // 忙碌
	INVISIABLE                   // 隐身
	OFFLINE                      // 离线
)

func (s Status) Int() int {
	return int(s)
}
