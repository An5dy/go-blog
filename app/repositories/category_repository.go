package repositories

import (
	"go-blog/app/models/category"
	"go-blog/pkg/database"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	All() (categories []category.IndexCategory)
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

func (cr *categoryRepository) Find(id int) {

}

func (cr *categoryRepository) Delete(id int) {

}
