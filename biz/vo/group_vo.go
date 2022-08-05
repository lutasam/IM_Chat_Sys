package vo

type GroupVO struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type GroupInSearchVO struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
