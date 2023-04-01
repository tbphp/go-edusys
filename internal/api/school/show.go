package school

import (
	"github.com/gin-gonic/gin"
	"github.com/tbphp/go-edusys/internal/e"
	"github.com/tbphp/go-edusys/internal/model"
	"github.com/tbphp/go-edusys/pkg/database"
)

func Show(c *gin.Context) model.School {
	var school model.School
	tx := database.DB.First(&school, c.Param("id"))
	if tx.Error != nil {
		panic(e.EmptyData)
	}

	return school
}
