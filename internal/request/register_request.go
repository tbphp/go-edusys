package request

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=50" label:"姓名"`
	Username string `json:"username" binding:"required,min=2,max=50,email" label:"用户名"`
	Password string `json:"password" binding:"required,min=6,max=32" label:"密码"`
}
