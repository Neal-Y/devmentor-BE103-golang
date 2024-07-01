package user

type Request struct {
	LineID      string `json:"line_id"`
	DisplayName string `json:"display_name" binding:"required"`
	Email       string `json:"email"`
	LineToken   string `json:"line_token"`
	Phone       string `json:"phone"`
	IsMember    bool   `json:"is_member"`
}
