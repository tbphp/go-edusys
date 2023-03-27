package migrate

import (
	"github.com/tbphp/go-edusys/internal/model"
	"github.com/tbphp/go-edusys/pkg/config"
	"github.com/tbphp/go-edusys/pkg/database"
)

var migrateModels = []any{
	model.School{},
	model.Teacher{},
	model.Student{},
}

func init() {
	if config.App.Mode != "debug" {
		return
	}

	for _, m := range migrateModels {
		err := database.DB.AutoMigrate(&m)
		if err != nil {
			panic(err)
		}
	}
}
