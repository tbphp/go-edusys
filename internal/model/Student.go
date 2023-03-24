package model

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/tbphp/go-edusys/internal/enum/identity"
	"github.com/tbphp/go-edusys/internal/request"
	"github.com/tbphp/go-edusys/pkg/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Student struct {
	AuthUser
	Name     string `json:"name" gorm:"type:varchar" binding:"required,min=2,max=50"`
	SchoolId int    `json:"school_id" gorm:"index:inx_school"`
	LineId   string `json:"line_id" gorm:"type:varchar(255);index:inx_line_id"`
	Times
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (t *Student) PayloadFunc() jwt.MapClaims {
	return jwt.MapClaims{
		"id":       t.ID,
		"identity": identity.Student,
	}
}

func (t *Student) Authenticator(req request.LoginRequest) (JwtUser, error) {
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
