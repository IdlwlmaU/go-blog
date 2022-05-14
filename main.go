package main

import (
	"fmt"
	"go-blog/dao/mysql"
	"go-blog/logger"
	"go-blog/router"
	"go-blog/settings"
	"go.uber.org/zap"
)

// Go Web开发较通用的脚手架模板

func main() {
	// 1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Println("init settings failed, err:", err)
		return
	}
	// 2.初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Println("init logger failed, err:", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	// 3.初始化MySQL连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Println("init mysql failed, err:", err)
		return
	}
	defer mysql.Close()
	// 4.注册路由
	r := router.SetUpRouter()
	// 5.启动服务
	addr := fmt.Sprintf(":%d", settings.Conf.Port)
	err := r.Run(addr)
	if err != nil {
		zap.L().Error("run server failed, err:", zap.Error(err))
		return
	}
}
