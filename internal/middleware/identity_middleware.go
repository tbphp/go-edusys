package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tbphp/go-edusys/internal/e"
	"github.com/tbphp/go-edusys/internal/model"
)

func IdentityMiddleware(ident int) gin.HandlerFunc {
	return func(c *gin.Context) {
		ui, ok := c.Get("auth")
		if !ok {
			panic(e.CodeError(e.Authorization))
		}

		if ui.(model.UserIdentity).Identity != ident {
			panic(e.CodeError(e.Authorization))
		}

		c.Next()
	}
}
