package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tbphp/go-edusys/internal/e"
	"github.com/tbphp/go-edusys/pkg/response"
	"net/http"
)

func CErrorRecoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if er, ok := err.(*e.CError); ok {
					result := response.NewResponse(er.Code(), er.Error(), map[string]any{})
					c.AbortWithStatusJSON(http.StatusOK, result)
				} else {
					panic(err)
				}
			}
		}()
		c.Next()
	}
}
