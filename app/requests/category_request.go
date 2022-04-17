package requests

import (
	"github.com/thedevsaddam/govalidator"
)

type StoreCategoryRequest struct {
	Ttile    string `valid:"title" json:"title"`
	ParentId string `valid:"parent_id" json:"parent_id"`
}

var _ FormRequest = (*StoreCategoryRequest)(nil)

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
