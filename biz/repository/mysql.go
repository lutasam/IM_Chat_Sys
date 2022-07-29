package repository

import (
	"github.com/lutasam/chat/biz/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(mysql.Open(utils.GetConfigString("mysql.dsn")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
