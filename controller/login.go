package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/logic"
	"go-blog/models"
	"go-blog/settings"
	"go.uber.org/zap"
	"net/http"
)

// LoginIndexHandler 登录界面加载
func LoginIndexHandler(c *gin.Context) {
	var obj = settings.Conf.Viewer
	c.HTML(http.StatusOK, "login.html", obj)
}

func LoginHandler(c *gin.Context) {
	params := new(models.Params)
	if err := c.ShouldBindJSON(params); err != nil {
		zap.L().Error("c.ShouldBindJSON() failed, err:", zap.Error(err))
		return
	}
	data, err := logic.Login(params.UserName, params.Passwd)
	if err != nil {
		c.JSON(http.StatusOK, &models.Result{
			Error: err.Error(),
			Data:  nil,
			Code:  -999,
		})
		return
	}
	models.ResponseSuccess(c, *data)
}
