package bootstrap

import (
	"go-blog/pkg/config"
	"go-blog/pkg/logger"
)

// SetupLogger 初始化 Logger
func SetupLogger() {
	logger.InitLogger(
		config.GetString("logging.filename"),
		config.GetInt("logging.max_size"),
		config.GetInt("logging.max_backup"),
		config.GetInt("logging.max_age"),
		config.GetBool("logging.compress"),
		config.GetString("logging.type"),
		config.GetString("logging.level"),
	)
}
