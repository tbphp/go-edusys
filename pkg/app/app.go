package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tbphp/go-edusys/internal/e"
	"github.com/tbphp/go-edusys/internal/middleware"
	"github.com/tbphp/go-edusys/internal/router"
	"github.com/tbphp/go-edusys/pkg/config"
	"github.com/tbphp/go-edusys/pkg/response"
)

var r *gin.Engine

func init() {
	gin.SetMode(config.App.Mode)
	r = gin.New()
	loadMiddlewares()
	registerRouters()
}

func Run() {
	_ = r.Run(fmt.Sprintf(
		"%s:%d",
		config.Server.HttpAddress,
		config.Server.HttpPort,
	))
}

func loadMiddlewares() {
	r.Use(
		gin.Logger(),
		gin.CustomRecovery(response.ErrorHandler()),
		middleware.CErrorRecoverMiddleware(),
	)
}

func registerRouters() {
	r.NoMethod(notFound())
	r.NoRoute(notFound())

	router.RegisterApiRouters(r)
}

func notFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		panic(e.CodeError(e.NotFound))
	}
}
