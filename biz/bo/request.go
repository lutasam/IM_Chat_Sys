package bo

type HelloRequest struct {
	Username string `json:"username"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
