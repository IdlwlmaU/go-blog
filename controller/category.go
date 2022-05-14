package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/logic"
	"net/http"
	"strconv"
)

func CategoryHandler(c *gin.Context) {
	// 1.获取参数并且进行参数校验
	cIdStr := c.Param("cId")
	cId, err := strconv.Atoi(cIdStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "路径不匹配",
		})
		return
	}
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	pageSize := 1000
	// 2.获取数据
	data := logic.GetPostByCategoryId(cId, page, pageSize)
	// 3.返回响应
	c.HTML(http.StatusOK, "category.html", data)
}
