package main

import (
	"go-blog/bootstrap"
	"go-blog/config"
	c "go-blog/pkg/config"

	"github.com/gin-gonic/gin"
)

func init() {
	config.Initialize()
}

func main() {
	// 初始化配置文件
	c.InitConfig()

	// 初始化日志
	bootstrap.SetupLogger()

	// 设置 gin 的运行模式，支持 debug, release, test
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	// 初始化数据库
	bootstrap.SetupDB()

	// 初始化路由
	bootstrap.SetupRoute(router)

	router.Run(":" + c.GetString("app.port"))
}
