package main

import (
	_ "github.com/tbphp/go-edusys/pkg/migrate"
	"github.com/tbphp/go-edusys/pkg/router"
)

func main() {
	router.Run()
}
