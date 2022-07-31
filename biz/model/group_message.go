package model

import (
	"gorm.io/gorm"
	"time"
)

type GroupMessage struct {
	ID        uint64         `gorm:"column:id"`
	GroupID   uint64         `gorm:"column:group_id"`
	UserID    uint64         `gorm:"column:user_id"`
	Content   string         `gorm:"column:content"`
	Name      string         `gorm:"column:name"`
	Avatar    string         `gorm:"column:avatar"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (GroupMessage) TableName() string {
	return "group_messages"
}
