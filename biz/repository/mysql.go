package repository

import (
	"fmt"
	"github.com/lutasam/chat/biz/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	DB     *gorm.DB
	DBOnce sync.Once
)

func GetDB() *gorm.DB {
	DBOnce.Do(func() {
		var err error
		DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
			utils.GetConfigResolve().GetConfigString("mysql.user"),
			utils.GetConfigResolve().GetConfigString("mysql.password"),
			utils.GetConfigResolve().GetConfigString("mysql.address"),
			utils.GetConfigResolve().GetConfigString("mysql.port"),
			utils.GetConfigResolve().GetConfigString("mysql.dbname"),
			utils.GetConfigResolve().GetConfigString("mysql.config"))), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	})
	return DB
}
