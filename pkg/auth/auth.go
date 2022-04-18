package auth

import (
	"errors"
	"go-blog/app/models/user"
	"go-blog/app/repositories"
	"go-blog/pkg/hash"
	"go-blog/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Attempt 尝试登录
func Attempt(username, password string) (user.User, error) {
	userModel, _ := repositories.NewUserRepository().FindByUsername(username)
	if userModel.ID == 0 {
		return user.User{}, errors.New("用户名错误。")
	}
	if !hash.Check(password, userModel.Password) {
		return user.User{}, errors.New("密码不正确。")
	}
	return userModel, nil
}

// SetUser 设置已认证的用户信息，用于后续获取用户信息
func SetUser(c *gin.Context, user user.User) {
	c.Set("current_user_id", user.ID)
	c.Set("current_username", user.Username)
	c.Set("current_user", user)
}

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}
	return userModel
}

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
