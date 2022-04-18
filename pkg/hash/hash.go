package hash

import (
	"go-blog/pkg/logger"

	"golang.org/x/crypto/bcrypt"
)

// Make 使用 bcrypt 对密码进行加密
func Make(password string) string {
	// GenerateFromPassword 的第二个参数是 cost 值。建议大于 12，数值越大耗费时间越长
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogIf(err)
	return string(bytes)
}

// Check 对比明文密码和数据库的哈希值
func Check(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// IsHashed 判断字符串是否是哈希过的数据
func IsHashed(str string) bool {
	return len(str) == 60
}
