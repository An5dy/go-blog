package middlewares

import (
	"fmt"
	"go-blog/app/models/user"
	"go-blog/app/repositories"
	"go-blog/pkg/auth"
	"go-blog/pkg/config"
	"go-blog/pkg/jwt"
	"go-blog/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		claims, err := jwt.NewJWT().ParserToken(c)

		// JWT 解析失败，有错误发生
		if err != nil {
			response.Abort401(c, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}

		// JWT 解析成功，设置用户信息
		userModel := user.User{}
		repositories.NewUserRepository().Find(&userModel, claims.UserID)
		if userModel.ID == 0 {
			response.Abort401(c, "查找的用户不存在。")
			return
		}

		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		auth.SetUser(c, userModel)
	}
}
