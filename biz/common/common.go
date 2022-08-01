package common

const ISSUER = "LUTASAM"
const JWTSECRETSALT = "LUTASAM"
const PASSWORDSALT = "astaxie12798akljzmknm.ahkjkljl;k"
const EXPIRETIME = 86400000
const CONFIGFILEPATH = "/root/go/src/chat/conf/config.yml"

const (
	STATUSOKCODE    = 200
	CLIENTERRORCODE = 400
	SERVERERRORCODE = 500
)

const (
	STATUSOKMSG    = "OK"
	CLIENTERRORMSG = "400 CLIENT ERROR"
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
