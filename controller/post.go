package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/logic"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

func PostDetailHandler(c *gin.Context) {
	// 获取参数
	pIdStr := c.Param("pid")
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		zap.L().Error("strconv.Atoi() failed, err:", zap.Error(err))
		return
	}
	// 处理业务
	data := logic.GetPostDetail(pId)
	// 返回响应
	c.HTML(http.StatusOK, "detail.html", data)
}
