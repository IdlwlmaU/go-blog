package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/logic"
	"net/http"
	"strconv"
)

func IndexHandler(c *gin.Context) {
	// 1.获取参数并进行参数校验
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	pageSize := 10
	// 2.获取数据
	hr := logic.GetIndexInfo(page, pageSize)
	// 3.返回响应
	c.HTML(http.StatusOK, "index.html", hr)
}
