package school

import (
	"github.com/gin-gonic/gin"
	"github.com/tbphp/go-edusys/internal/model"
	"github.com/tbphp/go-edusys/pkg/database"
	"github.com/tbphp/go-edusys/pkg/paginator"
)

func Index(c *gin.Context) *paginator.Page {
	var schools []model.School
	tx := database.DB.Order("id desc").Find(&schools)

	return paginator.NewPage(c, tx)
}
