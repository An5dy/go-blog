package routes

import (
	"go-blog/app/http/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册 API 路由
func RegisterAPIRoutes(router *gin.Engine) {
	v1 := router.Group("v1")
	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"hello": "world!",
			})
		})
		cg := v1.Group("category")
		{
			cc := controllers.NewCategoryController()
			cg.GET("", cc.Index)
			cg.POST("", cc.Store)
			cg.PUT("/:id", cc.Update)
			cg.DELETE("/:id", cc.Delete)
		}
	}
}
