package common

const ISSUER = "LUTASAM"                                // jwt issuer
const PASSWORDSALT = "astaxie12798akljzmknm.ahkjkljl;k" // use only for password encryption
const OTHERSECRETSALT = "9871267812345mn812345xyz"      // user for other encryption
const EXPIRETIME = 86400000                             // jwt expiration time. 1 day's second
const MAXMESSAGENUM = 99                                // max massage num on UI

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

type Status int // Status user status

const (
	ONLINE Status = iota + 1
	BUSY
	INVISIABLE
	OFFLINE
)

func (s Status) Int() int {
	return int(s)
}
