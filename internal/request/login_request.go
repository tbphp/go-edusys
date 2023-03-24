package request

type LoginRequest struct {
	Identity int    `json:"identity" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
