package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"hello": "world!",
		})
	})
	router.Run()
}
