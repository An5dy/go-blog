package app

import (
	"go-blog/pkg/config"
	"time"
)

// IsLocal 是否本地环境
func IsLocal() bool {
	return config.GetString("app.env") == "local"
}

// IsProduction 是否生成环境
func IsProduction() bool {
	return config.GetString("app.env") == "production"
}

// IsTesting 是否测试环境
func IsTesting() bool {
	return config.GetString("app.env") == "testing"
}

// IsDebug 是否调试
func IsDebug() bool {
	return config.GetBool("app.debug")
}

// TimenowInTimezone 获取当前时间，支持时区
func TimenowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.GetString("app.timezone"))
	return time.Now().In(chinaTimezone)
}
