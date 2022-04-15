package category

import (
	"go-blog/app/models"
)

type Category struct {
	models.ID
	Title    string `gorm:"column:title;type:varchar(255);not null;comment:分类标题;" json:"title"`
	ParentId uint64 `gorm:"column:parent_id;not null;default:0;index:idx_pid;comment:父级 ID;" json:"parent_id"`
	Level    uint8  `gorm:"column:level;not null;default:0;comment:层级;" json:"level,omitempty"`
	Path     string `gorm:"column:path;not null;default:'-';comment:层级路径" json:"path,omitempty"`
	models.Timestamps
}

type IndexCategory struct {
	ID       uint64 `json:"id"`
	Title    string `json:"title"`
	ParentId uint64 `json:"parent_id"`
}

type TreeNode struct {
	ID       uint64     `json:"id"`
	Title    string     `json:"title"`
	Children []TreeNode `json:"children,omitempty"`
}
