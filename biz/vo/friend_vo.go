package vo

type FriendVO struct {
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	MessageNum  int    `json:"message_num"`
	LastMessage string `json:"last_message"`
}
