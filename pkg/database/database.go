package database

import (
	"fmt"
	"github.com/tbphp/go-edusys/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		config.Database.Host,
		config.Database.User,
		config.Database.Password,
		config.Database.DbName,
		config.Database.Port,
		config.Database.TimeZone,
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
