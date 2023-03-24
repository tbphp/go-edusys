package paginator

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type Page struct {
	List        any   `json:"list"`
	Total       int64 `json:"total"`
	CurrentPage int   `json:"current_page"`
	LastPage    int   `json:"last_page"`
	PerPage     int   `json:"per_page"`
}

func NewPage(c *gin.Context, tx *gorm.DB) *Page {
	p := Page{}
	total := tx.RowsAffected
	perStr := c.DefaultQuery("per_page", "0")
	perPage, err := strconv.Atoi(perStr)
	if err != nil {
		perPage = 0
	}

	pagintor := NewPaginator(c.Request, perPage, total)
	p.Total = pagintor.Nums()
	p.CurrentPage = pagintor.Page()
	p.LastPage = pagintor.PageNums()
	p.PerPage = pagintor.PerPageNums

	tx.Offset(pagintor.Offset()).Limit(pagintor.PerPageNums).Find(tx.Statement.Dest)

	p.List = tx.Statement.Dest

	return &p
}
