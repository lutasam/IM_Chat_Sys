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
	USERNAMEDOESNOTEXIST = Error{
		ErrorCode:   10004,
		ErrorString: "username does not exist. please check",
	}
	PASSWORDISERROR = Error{
		ErrorCode:   10005,
		ErrorString: "password is incorrect. please try again",
	}
)
