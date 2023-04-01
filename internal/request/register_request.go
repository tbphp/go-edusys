package request

type RegisterRequest struct {
	Name     string `json:"name" binding:"required" label:"姓名"`
	Username string `json:"username" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required" label:"密码"`
}
