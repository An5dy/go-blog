package controllers

import (
	"go-blog/app/requests"
	"go-blog/app/services"
	"go-blog/pkg/request"
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
	categories, err := cc.categoryService.GetCategoryTree()
	if err != nil {
		response.Error(c, err)
	} else {
		response.Data(c, categories)
	}
}

func (cc *CategoryController) Store(c *gin.Context) {
	data := &requests.StoreCategoryRequest{}
	if ok := request.Validate(c, data); !ok {
		return
	}
	if _, ok, err := cc.categoryService.StoreCategory(data); !ok {
		response.Error(c, err)
	} else {
		response.Created(c)
	}
}

func (cc *CategoryController) Update(c *gin.Context) {
	data := &requests.UpdateCategoryRequest{}
	if ok := request.Validate(c, data); !ok {
		return
	}
}

func (cc *CategoryController) Delete(c *gin.Context) {
	if ok, err := cc.categoryService.DeleteCategoryById(c.Param("id")); !ok {
		response.Error(c, err)
	} else {
		response.Succeed(c)
	}
}
