package model

import (
	"gorm.io/gorm"
	"time"
)

type UserMessage struct {
	ID            uint64         `gorm:"column:id"`
	SendUserID    uint64         `gorm:"column:send_user_id"`
	ReceiveUserID uint64         `gorm:"column:receive_user_id"`
	Content       string         `gorm:"column:content"`
	Name          string         `gorm:"column:name"`
	Avatar        string         `gorm:"column:avatar"`
	CreatedAt     time.Time      `gorm:"column:created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (UserMessage) TableName() string {
	return "user_messages"
}
