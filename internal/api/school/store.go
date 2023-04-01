package school

import (
	"github.com/gin-gonic/gin"
	"github.com/tbphp/go-edusys/internal/e"
	"github.com/tbphp/go-edusys/internal/model"
	"github.com/tbphp/go-edusys/pkg/database"
)

func Store(c *gin.Context) {
	var school model.School
	if err := c.ShouldBindJSON(&school); err != nil {
		panic(err)
	}

	tx := database.DB.Create(&school)
	if tx.Error != nil {
		panic(e.Exception("创建失败"))
	}
}
