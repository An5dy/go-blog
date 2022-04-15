package services

import (
	"go-blog/app/models/category"
	"go-blog/app/repositories"
)

type CategoryService interface {
	GetCategoryTree() []category.TreeNode
}

type categoryService struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService() CategoryService {
	return &categoryService{
		categoryRepository: repositories.NewCategoryRepository(),
	}
}

func (cs *categoryService) GetCategoryTree() []category.TreeNode {
	categories := cs.categoryRepository.All()
	return cs.categoryRepository.GetTree(categories, 0)
}
