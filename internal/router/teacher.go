package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tbphp/go-edusys/internal/api/school"
)

func registerTeacherRouters(r *gin.RouterGroup) {
	// 学校接口
	r.GET("schools", h(school.Index))
	r.GET("schools/:id", h(school.Show))
	r.POST("schools", h(school.Store))
}
