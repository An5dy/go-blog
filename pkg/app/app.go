package app

import "go-blog/pkg/config"

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
