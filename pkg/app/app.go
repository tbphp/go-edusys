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
	gin.SetMode(config.App.Mode)
	InitLog()
	Migrate()

	r = gin.New()
	loadMiddlewares()
	registerRouters()
	InitSwagger()
}

func Run() {
	addr := fmt.Sprintf(
		"%s:%d",
		config.Server.HttpAddress,
		config.Server.HttpPort,
	)
	if config.Server.TLS {
		_ = r.RunTLS(addr, config.Server.TLSCertFile, config.Server.TLSKeyFile)
	} else {
		_ = r.Run(addr)
	}
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
		panic(e.NotFound)
	}
}
