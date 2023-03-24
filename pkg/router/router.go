package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tbphp/go-edusys/internal/e"
	"github.com/tbphp/go-edusys/internal/router"
	"github.com/tbphp/go-edusys/pkg/config"
	"github.com/tbphp/go-edusys/pkg/response"
)

var r *gin.Engine

func init() {
	gin.SetMode(config.App.Mode)
	r = gin.New()
	r.NoMethod(notFound)
	r.NoRoute(notFound)
	r.Use(gin.Logger(), gin.CustomRecovery(response.ErrorHandler()))
}

func register() {
	router.RegisterApiRouters(r)
}

func notFound(c *gin.Context) {
	panic(e.CodeError(e.NotFound))
}

func Run() {
	register()

	_ = r.Run(fmt.Sprintf(
		"%s:%d",
		config.Server.HttpAddress,
		config.Server.HttpPort,
	))
}
