package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tbphp/go-edusys/internal/e"
)

type PublicApi struct {
}

func NewPublicApi() *PublicApi {
	return &PublicApi{}
}

func (t *PublicApi) Index(c *gin.Context) {
	c.JSON(e.Ok, gin.H{
		"pong": "ok",
	})
}
