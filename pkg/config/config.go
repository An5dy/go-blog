package config

import (
	"go-blog/pkg/helpers"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

// viper 实例
var instance *viper.Viper

// ConfigFunc 动态加载配置信息
type ConfigFunc func() map[string]interface{}

// ConfigFuncs 先加载到此数组，loadConfig 再动态生成配置信息
var ConfigFuncs map[string]ConfigFunc

func init() {
	// 1. 初始化
	instance = viper.New()
	// 2. 配置类型，支持 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	instance.SetConfigType("env")
	// 3. 环境变量配置文件查找的路径，相对于 main.go
	instance.AddConfigPath(".")
	// 4. 设置环境变量前缀，用以区分 Go 的系统环境变量
	instance.SetEnvPrefix("appenv")
	// 5. 读取环境变量（支持 flags）
	instance.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

// InitConfig 初始化配置信息，完成对环境变量以及 config 信息的加载
func InitConfig() {
	loadEnv()
	loadConfig()
}

// loadEnv 加载环境变量
func loadEnv() {
	// 加载 env
	instance.SetConfigName(".env")
	if err := instance.ReadInConfig(); err != nil {
		panic(err)
	}

	// 监控 .env 文件，变更时重新加载
	instance.WatchConfig()
}

// loadConfig 注册配置信息
func loadConfig() {
	for key, fn := range ConfigFuncs {
		instance.Set(key, fn())
	}
}

// Env 读取环境变量，支持默认值
func Env(key string, defaultVal ...interface{}) interface{} {
	return Get(key, defaultVal...)
}

// Add 新增配置项
func Add(key string, fn ConfigFunc) {
	ConfigFuncs[key] = fn
}

// Get 获取配置项，允许使用点式获取，如：app.name，支持默认值
func Get(key string, defaultVal ...interface{}) interface{} {
	if !instance.IsSet(key) || helpers.Empty(instance.Get(key)) {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return nil
	}
	return instance.Get(key)
}

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(Get(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(Get(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(Get(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(Get(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(Get(path, defaultValue...))
}
