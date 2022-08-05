package vo

import "time"

type UserMessagesVO struct {
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
	Content string `json:"content"`
}

type GroupMessagesVO struct {
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
	Content string `json:"content"`
}

type MessageTipVO struct {
	Name        string    `json:"name"`
	Avatar      string    `json:"avatar"`
	MessageNum  int       `json:"message_num"`
	LastMessage string    `json:"last_message"`
	CreatedAt   time.Time `json:"created_at"`
}
