package database

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var (
	DB    *gorm.DB
	SQLDB *sql.DB
)

// Connect 连接数据库
func Connect(dbConfig gorm.Dialector, _logger gormlogger.Interface) {
	var err error
	// 使用 gorm.Open 连接数据库
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger:                                   _logger, // 设置日志
		DisableForeignKeyConstraintWhenMigrating: true,    // 禁用为关联创建外键约束
	})

	// 处理错误
	if err != nil {
		fmt.Println(err.Error())
	}

	// 获取底层的 SQLDB
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}
