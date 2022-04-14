package article

import "go-blog/app/models"

type Article struct {
	models.ID
	IsPublished bool   `gorm:"column:is_published;not null;default:0;comment:是否发布;" json:"is_published"`
	CategoryId  uint64 `gorm:"column:category_id;not null;comment:分类ID" json:"category_id"`
	Title       string `gorm:"column:title;type:varchar(255);not null;index;comment:文章标题;" json:"title,omitempty"`
	Sketch      string `gorm:"column:sketch;type:varchar(255);not null;comment:文章简述;" json:"sketch,omitempty"`
	Content     string `gorm:"column:content;type:text;not null;comment:文章内容;" json:"content,omitempty"`
	Skims       uint32 `gorm:"column:skims;not null;default:0;comment:浏览量;" json:"skims"`
	Likes       uint32 `gorm:"column:likes;not null;default:0;comment:点赞量;" json:"likes"`
	Comments    uint32 `gorm:"column:comments;not null;default:0;comment:评论量;" json:"comments"`
	models.Timestamps
	models.DeletedAt
}
