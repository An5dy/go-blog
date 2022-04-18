package requests

import (
	"github.com/thedevsaddam/govalidator"
)

// 验证新增分类请求数据
type StoreCategoryRequest struct {
	Ttile    string `valid:"title" json:"title"`
	ParentId string `valid:"parent_id" json:"parent_id"`
}

func (r *StoreCategoryRequest) Rules() govalidator.MapData {
	return govalidator.MapData{
		"title":     []string{"required"},
		"parent_id": []string{"required", "numeric"},
	}
}

func (r *StoreCategoryRequest) Messages() govalidator.MapData {
	return govalidator.MapData{
		"title": []string{
			"required:分类名称不能为空",
		},
		"parent_id": []string{
			"required:父级分类不能为空",
			"numeric:父级分类必须为数字",
		},
	}
}

type UpdateCategoryRequest struct {
	Title string `valid:"title" json:"title"`
}

func (r *UpdateCategoryRequest) Rules() govalidator.MapData {
	return govalidator.MapData{
		"title": []string{"required"},
	}
}

func (r *UpdateCategoryRequest) Messages() govalidator.MapData {
	return govalidator.MapData{
		"title": []string{
			"required:分类名称不能为空",
		},
	}
}
