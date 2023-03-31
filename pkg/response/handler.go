package response

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/tbphp/go-edusys/internal/e"
	"net/http"
)

func ErrorHandler() gin.RecoveryFunc {
	return func(c *gin.Context, err any) {
		defaultError := e.Default
		result := NewResponse(defaultError.Code(), defaultError.Error(), map[string]any{})

		isCe := false
		switch v := err.(type) {
		case *e.CError:
			result.Code = v.Code()
			result.Msg = v.Error()
			isCe = true
			break
		case error:
			result.Msg = v.Error()
			break
		case string:
			result.Msg = v
			break
		}

		if !isCe {
			log.Error(err)
		}

		c.AbortWithStatusJSON(http.StatusOK, result)
	}
}
