package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tbphp/go-edusys/internal/api"
	"github.com/tbphp/go-edusys/pkg/database"
)

func registerTeacherRouters(r *gin.RouterGroup) {
	school := api.NewSchoolApi(database.DB)
	// 学校接口
	r.GET("schools", school.Index)
	r.GET("schools/:id", school.Show)
	r.POST("schools", school.Store)
}
