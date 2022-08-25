package common

const ISSUER = "LUTASAM"                                // jwt issuer
const PASSWORDSALT = "astaxie12798akljzmknm.ahkjkljl;k" // use only for password encryption
const OTHERSECRETSALT = "9871267812345mn812345xyz"      // user for other encryption
const EXPIRETIME = 86400000                             // jwt expiration time. 1 day's second
const MAXMESSAGENUM = 99                                // max massage num on UI
const DEFAULTAVATARURL = "http://baidu.com/test.png"
const DEFAULTDESCRIPTION = "This group of men doesn't say anything."
const DEFAULTSIGN = "This man doesn't say anything."
const DEFAULTNICKNAME = "SHAREUSER"

const (
	USERDETAILPREFIX = "user_detail_"
)

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

// Status user status
type Status int

const (
	ONLINE Status = iota + 1
	BUSY
	INVISIBLE
	OFFLINE
)

func ParseInt2Status(i int) Status {
	switch i {
	case 1:
		return ONLINE
	case 2:
		return BUSY
	case 3:
		return INVISIBLE
	case 4:
		return OFFLINE
	default:
		return ONLINE
	}
}

func ParseString2Status(s string) Status {
	switch s {
	case "ONLINE":
		return ONLINE
	case "BUSY":
		return BUSY
	case "INVISIBLE":
		return INVISIBLE
	case "OFFLINE":
		return OFFLINE
	default:
		return ONLINE
	}
}

func (s Status) Int() int {
	switch s {
	case ONLINE:
		return 1
	case BUSY:
		return 2
	case INVISIBLE:
		return 3
	case OFFLINE:
		return 4
	default:
		return 1
	}
}

func (s Status) String() string {
	switch s {
	case ONLINE:
		return "ONLINE"
	case BUSY:
		return "BUSY"
	case INVISIBLE:
		return "INVISIBLE"
	case OFFLINE:
		return "OFFLINE"
	default:
		return "ONLINE"
	}
}
