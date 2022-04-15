package controllers

import (
	"go-blog/app/services"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	articleService *services.ArticleService
}

func NewArticleController() *ArticleController {
	return &ArticleController{
		articleService: services.NewArticleService(),
	}
}

func (ac *ArticleController) Index(c *gin.Context) {

}

func (ac *ArticleController) Show(c *gin.Context) {

}

func (ac *ArticleController) Store(c *gin.Context) {

}

func (ac *ArticleController) Update(c *gin.Context) {

}

func (ac *ArticleController) Delete(c *gin.Context) {

}
