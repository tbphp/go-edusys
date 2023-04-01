package request

type LoginRequest struct {
	Identity int    `json:"identity" binding:"required" label:"身份"`
	Username string `json:"username" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required" label:"密码"`
}
