package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64         `gorm:"column:id"`
	Account   string         `gorm:"column:account"`
	Password  string         `gorm:"column:password"`
	NickName  string         `gorm:"column:nickname"`
	Avatar    string         `gorm:"column:avatar"`
	Sign      string         `gorm:"column:sign"`
	IP        string         `gorm:"column:ip"`
	Port      int            `gorm:"column:port"`
	Status    int            `gorm:"column:status"`
	Groups    []*Group       `gorm:"many2many:users_groups"`
	Friends   []*User        `gorm:"many2many:users_friends"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (User) TableName() string {
	return "users"
}
