package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/models"
	"go-blog/settings"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "dema",
			ViewCount:    10086,
			CreateAt:     "2022-05-14",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &models.HomeResponse{
		Viewer:    *settings.Conf.Viewer,
		Categorys: categorys,
		Posts:     posts,
		Total:     1,
		Page:      1,
		Pages:     []int{1},
		PageEnd:   true,
	}
	c.HTML(http.StatusOK, "index.html", hr)
}
