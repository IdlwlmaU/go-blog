package models

import (
	"go-blog/settings"
)

type HomeResponse struct {
	*settings.Viewer
	Categorys []Category
	Posts     []PostMore
	Total     int
	Page      int
	Pages     []int
	PageEnd   bool
}
