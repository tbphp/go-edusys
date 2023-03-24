package model

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/tbphp/go-edusys/internal/enum/identity"
	"github.com/tbphp/go-edusys/internal/request"
	"github.com/tbphp/go-edusys/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

type Teacher struct {
	AuthUser
	Name   string `json:"name" gorm:"type:varchar" binding:"required,min=2,max=50"`
	LineId string `json:"line_id" gorm:"type:varchar(255);unique:unq_line_id"`
	Times
}

func (t *Teacher) PayloadFunc() jwt.MapClaims {
	return jwt.MapClaims{
		"id":       t.ID,
		"identity": identity.Teacher,
	}
}

func (t *Teacher) Authenticator(req request.LoginRequest) (JwtUser, error) {
	t.Username = req.Username
	tx := database.DB.Where(t).First(t)
	if tx.Error != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	err := bcrypt.CompareHashAndPassword([]byte(t.Password), []byte(req.Password))
	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	return t, nil
}
