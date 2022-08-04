package vo

type FriendVO struct {
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	MessageNum  int    `json:"message_num"`
	LastMessage string `json:"last_message"`
}

type FriendInSearchVO struct {
	ID       string `json:"id"`
	Account  string `json:"account"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}
