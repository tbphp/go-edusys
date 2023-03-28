package main

import (
	"github.com/tbphp/go-edusys/pkg/gin"
	_ "github.com/tbphp/go-edusys/pkg/logrus"
	_ "github.com/tbphp/go-edusys/pkg/migrate"
)

func main() {
	gin.Run()
}
