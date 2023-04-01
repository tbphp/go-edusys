package school

import (
	"github.com/tbphp/go-edusys/internal/e"
	"github.com/tbphp/go-edusys/internal/model"
	"github.com/tbphp/go-edusys/pkg/database"
)

func Store(school model.School) {
	tx := database.DB.Create(&school)
	if tx.Error != nil {
		panic(e.Exception("创建失败"))
	}
}
