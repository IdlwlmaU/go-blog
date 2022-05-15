package router

import (
	"github.com/gin-gonic/gin"
	"go-blog/controller"
	"go-blog/logger"
	"html/template"
)

func SetUpRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 告诉gin框架模板文件需要的静态文件在哪
	r.Static("/resource/", "./public/resource/")
	// 对应模板中的函数
	r.SetFuncMap(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date, "dateDay": DateDay})
	// 告诉gin框架去哪个文件夹去找模板文件
	r.LoadHTMLGlob("./template/*")

	r.GET("/", controller.IndexHandler)

	r.GET("/c/:cId", controller.CategoryHandler)

	r.GET("/login", controller.LoginIndexHandler)

	r.GET("/p/:pid", controller.PostDetailHandler)

	r.POST("/api/v1/login", controller.LoginHandler)

	return r
}
