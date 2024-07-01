package admin

type UpdateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
