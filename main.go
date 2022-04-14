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
	c.InitConfig()

	router := gin.New()

	bootstrap.SetupRoute(router)

	router.Run(":" + c.GetString("app.port"))
}
