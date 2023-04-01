package auth

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/tbphp/go-edusys/internal/e"
	"github.com/tbphp/go-edusys/internal/enum/identity"
	"github.com/tbphp/go-edusys/internal/middleware"
	"github.com/tbphp/go-edusys/internal/model"
	"github.com/tbphp/go-edusys/internal/request"
	"github.com/tbphp/go-edusys/pkg/config"
	"time"
)

var Auth *jwt.GinJWTMiddleware

func init() {
	var err error
	Auth, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "authuser",
		Key:         []byte(config.Auth.Key),
		Timeout:     time.Hour * 24 * 365 * 10,
		IdentityKey: "auth",
		PayloadFunc: func(data any) jwt.MapClaims {
			if v, ok := data.(model.JwtUser); ok {
				return v.PayloadFunc()
			}

			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (any, error) {
			var req request.LoginRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				panic(err)
			}

			var user model.JwtUser
			if req.Identity == identity.Teacher {
				user = &model.Teacher{}
			} else {
				user = &model.Student{}
			}

			return user.Authenticator(req)
		},
		LoginResponse: func(c *gin.Context, code int, token string, time time.Time) {
			c.Set(middleware.HandlerResultKey, gin.H{
				"token_type":   "bearer",
				"access_token": token,
			})
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			panic(e.Unauthorized(message))
		},
		IdentityHandler: func(c *gin.Context) any {
			claims := jwt.ExtractClaims(c)
			authUser := model.AuthUser{ID: uint(claims["id"].(float64))}

			var user model.JwtUser
			var it int
			if int(claims["identity"].(float64)) == identity.Teacher {
				it = identity.Teacher
				user = &model.Teacher{AuthUser: authUser}
			} else {
				it = identity.Student
				user = &model.Student{AuthUser: authUser}
			}
			return model.UserIdentity{
				Identity: it,
				ID:       authUser.ID,
				User:     user,
			}
		},
	})

	if err != nil {
		panic(fmt.Sprintf("JWT初始化失败: %v", err))
	}
}
