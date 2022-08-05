package vo

type FriendVO struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type FriendInSearchVO struct {
	ID       string `json:"id"`
	Account  string `json:"account"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}
