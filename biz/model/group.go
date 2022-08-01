package model

import (
	"time"

	"gorm.io/gorm"
)

type Group struct {
	ID        uint64         `gorm:"column:id"`
	Name      string         `gorm:"column:name"`
	Describe  string         `gorm:"column:describe"`
	Avatar    string         `gorm:"column:avatar"`
	AdminID   uint64         `gorm:"column:admin_id"`
	User      []*User        `gorm:"many2many:users_groups"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Group) TableName() string {
	return "groups"
}
