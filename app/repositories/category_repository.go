package repositories

import (
	"go-blog/app/models/category"
	"go-blog/pkg/database"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	All() (categories []category.IndexCategory)
	GetTree(categories []category.IndexCategory, parentId uint64) []category.TreeNode
	Find(id int)
	Delete(id int)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{
		db: database.DB,
	}
}

func (cr *categoryRepository) All() (categories []category.IndexCategory) {
	cr.db.Model(&category.Category{}).Find(&categories)
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

func (cr *categoryRepository) Find(id int) {

}

func (cr *categoryRepository) Delete(id int) {

}
