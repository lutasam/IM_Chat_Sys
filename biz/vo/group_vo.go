package vo

type GroupWithMessageVO struct {
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	MessageNum  int    `json:"message_num"`
	LastMessage string `json:"last_message"`
}
