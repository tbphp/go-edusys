package response

import (
	"github.com/gin-gonic/gin"
	"github.com/tbphp/go-edusys/internal/e"
	"net/http"
)

func ErrorHandler() gin.RecoveryFunc {
	return func(c *gin.Context, err any) {
		result := NewResponse(e.Exception, e.GetMsg(e.Exception), map[string]any{})

		switch v := err.(type) {
		case error:
			result.Msg = v.Error()
			break
		case string:
			result.Msg = v
			break
		}

		c.AbortWithStatusJSON(http.StatusOK, result)
	}
}
