package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tbphp/go-edusys/internal/api"
	"github.com/tbphp/go-edusys/internal/middleware"
	"github.com/tbphp/go-edusys/pkg/database"
)

func RegisterApiRouters(r *gin.Engine) {
	g := r.Group("/api", middleware.ResultMiddleware())
	{
		public := api.NewPublicApi()
		g.GET("/ping", public.Index)

		// 学校接口
		school := api.NewSchoolApi(database.DB)
		g.GET("/schools", school.Index)
		g.GET("/schools/:id", school.Show)
		g.POST("/schools", school.Store)
	}
}
