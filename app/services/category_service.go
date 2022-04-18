package services

import (
	"go-blog/app/models/category"
	"go-blog/app/repositories"
	"go-blog/app/requests"

	"github.com/spf13/cast"
)

type CategoryService interface {
	GetCategoryTree() (categoryTree []category.TreeNode, err error)
	StoreCategory(*requests.StoreCategoryRequest) (category category.Category, ok bool, err error)
	UpdateCategoryById(id string) (ok bool, err error)
	DeleteCategoryById(id string) (ok bool, err error)
}

type categoryService struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService() CategoryService {
	return &categoryService{
		categoryRepository: repositories.NewCategoryRepository(),
	}
}

func (cs *categoryService) GetCategoryTree() (categoryTree []category.TreeNode, err error) {
	categories, result := cs.categoryRepository.All()
	if result.Error != nil {
		return categoryTree, result.Error
	}
	return cs.categoryRepository.GetTree(categories, 0), nil
}

func (cs *categoryService) StoreCategory(data *requests.StoreCategoryRequest) (category category.Category, ok bool, err error) {
	category.Title = data.Ttile
	category.ParentId = cast.ToUint64(data.ParentId)
	result := cs.categoryRepository.Create(&category)
	if result.Error != nil {
		return category, false, result.Error
	}
	return category, true, result.Error
}

func (cs *categoryService) UpdateCategoryById(id string) (ok bool, err error) {
	// 查询分类
	categoryModel := &category.Category{}
	result := cs.categoryRepository.Find(categoryModel, cast.ToUint64(id))
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (cs *categoryService) DeleteCategoryById(id string) (ok bool, err error) {
	// 查询分类
	categoryModel := &category.Category{}
	result := cs.categoryRepository.Find(categoryModel, cast.ToUint64(id))
	if result.Error != nil {
		return false, result.Error
	}
	// 删除分类
	result = cs.categoryRepository.Delete(categoryModel)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
