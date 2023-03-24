package model

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/tbphp/go-edusys/internal/request"
)

type JwtUser interface {
	PayloadFunc() jwt.MapClaims
	Authenticator(req request.LoginRequest) (JwtUser, error)
}

type UserIdentity struct {
	Identity int
	ID       uint
	User     JwtUser
}

type AuthUser struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"type:varchar;unique:unq_username" binding:"required,min=2,max=50"`
	Password string `json:"-" gorm:"type:varchar(255)" binding:"required,min=6,max=50"`
}
