package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/tbphp/go-edusys/internal/e"
	"github.com/tbphp/go-edusys/internal/router"
	"github.com/tbphp/go-edusys/pkg/config"
	"github.com/tbphp/go-edusys/pkg/response"
	ginlogrus "github.com/toorop/gin-logrus"
)

var r *gin.Engine

func init() {
	InitLog()
	Migrate()

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
	if !gin.IsDebugging() {
		r.Use(ginlogrus.Logger(log.StandardLogger()))
	}
	r.Use(gin.CustomRecoveryWithWriter(nil, response.ErrorHandler()))
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
