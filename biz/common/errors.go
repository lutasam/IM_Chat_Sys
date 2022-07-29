package common

type Error struct {
	error
	ErrorString string
	ErrorCode   int
}

func (e Error) Error() string {
	return e.ErrorString
}

func (e Error) Code() int {
	return e.ErrorCode
}

var (
	UNKNOWNERROR = Error{
		ErrorCode:   -1,
		ErrorString: "unknown error. maybe server is error. please wait for sometime",
	}
	USERINPUTERROR = Error{
		ErrorCode:   10001,
		ErrorString: "please check your input, there is something wrong",
	}
	HAVENOPERMISSION = Error{
		ErrorCode:   10002,
		ErrorString: "you have no access to this operation",
	}
	DATABASEERROR = Error{
		ErrorCode:   10003,
		ErrorString: "server's database has some error, please try again later",
	}
	USERDOESNOTEXIST = Error{
		ErrorCode:   10004,
		ErrorString: "user does not exist. please check",
	}
	PASSWORDISERROR = Error{
		ErrorCode:   10005,
		ErrorString: "password is incorrect. please try again",
	}
	USERNOTLOGIN = Error{
		ErrorCode:   10006,
		ErrorString: "you do not login. please login",
	}
	EXCEEDTIMELIMIT = Error{
		ErrorCode:   10007,
		ErrorString: "your token has no time. please login again",
	}
)
