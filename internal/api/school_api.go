package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tbphp/go-edusys/internal/e"
	"github.com/tbphp/go-edusys/internal/middleware"
	"github.com/tbphp/go-edusys/internal/model"
	"github.com/tbphp/go-edusys/pkg/paginator"
	"gorm.io/gorm"
)

type SchoolApi struct {
	db *gorm.DB
}

func NewSchoolApi(db *gorm.DB) *SchoolApi {
	return &SchoolApi{db: db}
}

func (t *SchoolApi) Index(c *gin.Context) {
	var schools []model.School
	tx := t.db.Order("id desc").Find(&schools)

	result := paginator.NewPage(c, tx)

	c.Set(middleware.HandlerResultKey, result)
}

func (t *SchoolApi) Show(c *gin.Context) {
	var school model.School
	tx := t.db.First(&school, c.Param("id"))
	if tx.Error != nil {
		panic(e.EmptyDataError())
	}

	c.Set(middleware.HandlerResultKey, school)
}

func (t *SchoolApi) Store(c *gin.Context) {
	var school model.School
	if err := c.ShouldBindJSON(&school); err != nil {
		panic(err)
	}

	tx := t.db.Create(&school)
	if tx.Error != nil {
		panic(tx.Error)
	}
}
