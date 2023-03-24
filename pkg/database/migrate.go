package database

import (
	"github.com/tbphp/go-edusys/internal/model"
	"gorm.io/gorm"
)

var migrateModels = []any{
	model.School{},
}

func AutoMigrate(db *gorm.DB) {
	for _, m := range migrateModels {
		err := db.AutoMigrate(&m)
		if err != nil {
			panic(err)
		}
	}
}
