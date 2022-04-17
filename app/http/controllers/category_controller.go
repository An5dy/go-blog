package controllers

import (
	"go-blog/app/services"
	"go-blog/pkg/response"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService services.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		categoryService: services.NewCategoryService(),
	}
}

func (cc *CategoryController) Index(c *gin.Context) {
	response.Data(c, cc.categoryService.GetCategoryTree())
}

func (cc *CategoryController) Store(c *gin.Context) {

}

func (cc *CategoryController) Update(c *gin.Context) {

}

func (cc *CategoryController) Delete(c *gin.Context) {

}
