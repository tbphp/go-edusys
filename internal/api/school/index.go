package school

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/tbphp/go-edusys/internal/middleware"
	"github.com/tbphp/go-edusys/internal/model"
	"github.com/tbphp/go-edusys/pkg/database"
	"github.com/tbphp/go-edusys/pkg/paginator"
)

func Index(c *gin.Context) {
	fmt.Println(jwt.ExtractClaims(c))
	var schools []model.School
	tx := database.DB.Order("id desc").Find(&schools)

	result := paginator.NewPage(c, tx)

	c.Set(middleware.HandlerResultKey, result)
}
