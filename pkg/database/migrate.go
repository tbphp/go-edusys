package database

import (
	"gorm.io/gorm"
)

var migrateModels = []any{
	// model.School{},
	// model.Teacher{},
	// model.Student{},
}

func AutoMigrate(db *gorm.DB) {
	for _, m := range migrateModels {
		err := db.AutoMigrate(&m)
		if err != nil {
			panic(err)
		}
	}
}
