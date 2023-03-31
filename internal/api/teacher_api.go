package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tbphp/go-edusys/internal/e"
	"github.com/tbphp/go-edusys/internal/model"
	"github.com/tbphp/go-edusys/internal/request"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type TeacherApi struct {
	db *gorm.DB
}

func NewTeacherApi(db *gorm.DB) *TeacherApi {
	return &TeacherApi{db: db}
}

func (t *TeacherApi) Register(c *gin.Context) {
	req := request.RegisterRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(e.ValidationError("参数错误"))
	}

	// 生成密码
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(e.DefaultError("密码加密失败"))
	}

	teacher := model.Teacher{
		AuthUser: model.AuthUser{
			Username: req.Username,
			Password: string(passwordBytes),
		},
		Name: req.Name,
	}
	tx := t.db.Create(&teacher)
	if tx.Error != nil {
		panic(e.DefaultError("创建失败"))
	}
}
