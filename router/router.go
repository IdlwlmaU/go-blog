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
	home := "template/home.html"
	footer := "template/layout/footer.html"
	header := "template/layout/header.html"
	personal := "template/layout/personal.html"
	postList := "template/layout/post-list.html"
	pagination := "template/layout/pagination.html"
	r.LoadHTMLFiles("template/index.html", home, footer, header, personal, postList, pagination)
	r.GET("/", controller.IndexHandler)
	return r
}
