package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tbphp/go-edusys/internal/api/teacher"
	"github.com/tbphp/go-edusys/internal/enum/identity"
	"github.com/tbphp/go-edusys/internal/middleware"
	"github.com/tbphp/go-edusys/pkg/auth"
)

func RegisterApiRouters(r *gin.Engine) {
	g := r.Group("api", middleware.ResultMiddleware())
	{
		g.POST("login", auth.Auth.LoginHandler)

		// 教师注册
		g.POST("teacher/register", h(teacher.Register))

		// 教师
		teacherGroup := g.Group(
			"teacher",
			auth.Auth.MiddlewareFunc(),
			middleware.IdentityMiddleware(identity.Teacher),
		)
		registerTeacherRouters(teacherGroup)

		// 学生
		studentGroup := g.Group(
			"student",
			auth.Auth.MiddlewareFunc(),
			middleware.IdentityMiddleware(identity.Student),
		)
		registerStudentRouters(studentGroup)
	}
}
