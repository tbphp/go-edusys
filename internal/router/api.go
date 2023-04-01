package router

import (
	"github.com/gin-gonic/gin"
	teacher2 "github.com/tbphp/go-edusys/internal/api/teacher"
	"github.com/tbphp/go-edusys/internal/enum/identity"
	"github.com/tbphp/go-edusys/internal/middleware"
	"github.com/tbphp/go-edusys/pkg/auth"
)

func RegisterApiRouters(r *gin.Engine) {
	g := r.Group("api", middleware.ResultMiddleware())
	{
		g.POST("login", auth.Auth.LoginHandler)

		// 教师注册
		g.POST("teacher/register", teacher2.Register)

		// 教师
		teacher := g.Group(
			"teacher",
			auth.Auth.MiddlewareFunc(),
			middleware.IdentityMiddleware(identity.Teacher),
		)
		registerTeacherRouters(teacher)

		// 学生
		student := g.Group(
			"student",
			auth.Auth.MiddlewareFunc(),
			middleware.IdentityMiddleware(identity.Student),
		)
		registerStudentRouters(student)
	}
}
