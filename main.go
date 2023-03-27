package main

import (
	"github.com/tbphp/go-edusys/pkg/app"
	_ "github.com/tbphp/go-edusys/pkg/migrate"
)

func main() {
	app.Run()
}
