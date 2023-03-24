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
			return
		}

		uit := ui.(model.UserIdentity)
		if uit.Identity != ident {
			panic(e.CodeError(e.Authorization))
		}

		c.Next()
	}
}
