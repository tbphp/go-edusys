package response

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github.com/tbphp/go-edusys/internal/e"
	validator2 "github.com/tbphp/go-edusys/pkg/validator"
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
		case validator.ValidationErrors:
			isCe = true
			result.Code = 422
			result.Msg = v[0].Translate(validator2.Trans)
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
