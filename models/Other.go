package Models

import (
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pages, ok := r.URL.Query()["page"]
		var page int = 1
		if ok && len(pages[0]) > 0 {
			page, _ = strconv.Atoi(pages[0])
		}
		if page == 0 || !ok {
			page = 1
		}

		pageSize, _ := strconv.Atoi(r.URL.Query()["page_size"][0])
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
