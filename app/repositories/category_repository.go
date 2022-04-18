package repositories

import (
	"go-blog/app/models/category"
	"go-blog/pkg/database"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	All() (categories []category.IndexCategory, result *gorm.DB)                      // 获取所有的分类
	GetTree(categories []category.IndexCategory, parentId uint64) []category.TreeNode // 获取分类树
	Create(category *category.Category) (result *gorm.DB)                             // 创建分类
	Find(category *category.Category, id uint64) (result *gorm.DB)                    // 查询分类
	Update(category *category.Category) (result *gorm.DB)                             // 更新分类
	Delete(category *category.Category) (result *gorm.DB)                             // 删除分类
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{
		db: database.DB,
	}
}

func (cr *categoryRepository) All() (categories []category.IndexCategory, result *gorm.DB) {
	result = cr.db.Model(&category.Category{}).Find(&categories)
	return
}

func (cr *categoryRepository) GetTree(categories []category.IndexCategory, parentId uint64) []category.TreeNode {
	var treeNodes []category.TreeNode
	for _, item := range categories {
		if item.ParentId == parentId {
			// 递归获取所有的子集
			child := cr.GetTree(categories, item.ID)
			node := category.TreeNode{
				ID:       item.ID,
				Title:    item.Title,
				Children: child,
			}
			treeNodes = append(treeNodes, node)
		}
	}
	return treeNodes
}

func (cr *categoryRepository) Create(category *category.Category) (result *gorm.DB) {
	return cr.db.Create(category)
}

func (cr *categoryRepository) Find(category *category.Category, id uint64) (result *gorm.DB) {
	return cr.db.First(category, id)
}

func (cr *categoryRepository) Update(category *category.Category) (result *gorm.DB) {
	return cr.db.Save(category)
}

func (cr *categoryRepository) Delete(category *category.Category) (result *gorm.DB) {
	return cr.db.Delete(category)
}
