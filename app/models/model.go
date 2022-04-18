package models

import (
	"database/sql"
	"time"
)

// PrimaryKey 表主键
type PrimaryKey struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;comment:主键" json:"id,omitempty"`
}

// Timestamps 表记录新增/更新时间
type Timestamps struct {
	CreatedAt time.Time `gorm:"column:created_at;comment:创建时间" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;comment:修改时间" json:"updated_at,omitempty"`
}

// SoftDelete 表记录软删除时间
type SoftDelete struct {
	DeletedAt sql.NullTime `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at,omitempty"`
}
