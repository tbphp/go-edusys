package migrate

import (
	"github.com/gin-gonic/gin"
	"github.com/tbphp/go-edusys/internal/model"
	"github.com/tbphp/go-edusys/pkg/database"
)

var migrateModels = []any{
	model.School{},
	model.Teacher{},
	model.Student{},
}

func init() {
	if !gin.IsDebugging() {
		return
	}

	for _, m := range migrateModels {
		err := database.DB.AutoMigrate(&m)
		if err != nil {
			panic(err)
		}
	}
}
