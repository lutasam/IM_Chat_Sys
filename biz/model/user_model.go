package model

type User struct {
	ID       uint64 `gorm:"column:id"`
	Account  string `gorm:"column:account"`
	Password string `gorm:"column:password"`
	NickName string `gorm:"column:nickname"`
	Avatar   string `gorm:"column:avatar"`
	Sign     string `gorm:"column:sign"`
	IP       string `gorm:"column:ip"`
	Port     int    `gorm:"column:port"`
	Status   int    `gorm:"column:status"`
}

func (User) TableName() string {
	return "users"
}
