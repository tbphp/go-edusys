package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tbphp/go-edusys/internal/router"
	"github.com/tbphp/go-edusys/pkg/config"
	"github.com/tbphp/go-edusys/pkg/response"
)

var r *gin.Engine

func init() {
	gin.SetMode(config.App.Mode)
	r = gin.New()
	r.Use(gin.Logger(), gin.CustomRecovery(response.ErrorHandler()))
}

func register() {
	router.RegisterApiRouters(r)
}

func Run() {
	register()

	_ = r.Run(fmt.Sprintf(
		"%s:%d",
		config.Server.HttpAddress,
		config.Server.HttpPort,
	))
}
