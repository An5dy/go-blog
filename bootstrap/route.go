package bootstrap

import (
	"go-blog/app/http/middlewares"
	"go-blog/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoute 初始化路由
func SetupRoute(router *gin.Engine) {
	registerGlobalMiddleWare(router)

	routes.RegisterAPIRoutes(router)

	setup404Handler(router)
}

// registerGlobalMiddleWare 注册全局中间件
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
	)
}

// setup404Handler 设置 404 请求
func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "404 Not Found!",
		})
	})
}
