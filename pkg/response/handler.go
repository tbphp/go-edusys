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
		case *e.CError:
			result.Code = v.Code()
			result.Msg = v.Error()
		case error:
			result.Msg = v.Error()
		}

		c.AbortWithStatusJSON(http.StatusOK, result)
	}
}
