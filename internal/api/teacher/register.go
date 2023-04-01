package teacher

import (
	"github.com/tbphp/go-edusys/internal/e"
	"github.com/tbphp/go-edusys/internal/model"
	"github.com/tbphp/go-edusys/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

type registerRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=50" label:"姓名"`
	Username string `json:"username" binding:"required,min=2,max=50,email" label:"用户名"`
	Password string `json:"password" binding:"required,min=6,max=32" label:"密码"`
}

func Register(req registerRequest) {
	// 生成密码
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(e.Exception("密码加密失败"))
	}

	teacher := model.Teacher{
		AuthUser: model.AuthUser{
			Username: req.Username,
			Password: string(passwordBytes),
		},
		Name: req.Name,
	}
	tx := database.DB.Create(&teacher)
	if tx.Error != nil {
		panic(e.Exception("创建失败"))
	}
}
