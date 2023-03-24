package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tbphp/go-edusys/internal/e"
	"github.com/tbphp/go-edusys/pkg/response"
	"net/http"
)

const HandlerResultKey = "r_handler_result"

func ResultMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 默认返回值
		result := response.NewResponse(e.Ok, e.GetMsg(e.Ok), map[string]any{})

		// 解析返回值
		data, ok := c.Get(HandlerResultKey)
		if ok && data != nil {
			result.Data = data
		}

		c.JSON(http.StatusOK, result)
	}
}
