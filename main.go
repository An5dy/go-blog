package main

import (
	"go-blog/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	bootstrap.SetupRoute(router)

	router.Run()
}
