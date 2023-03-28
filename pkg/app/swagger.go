package app

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tbphp/go-edusys/docs"
)

func InitSwagger() {
	if !gin.IsDebugging() {
		return
	}
	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
