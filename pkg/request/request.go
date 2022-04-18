package request

import (
	"go-blog/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type FormRequest interface {
	Rules() govalidator.MapData    // 自定义表单验证规则
	Messages() govalidator.MapData // 自定义错误验证信息
}

// Validate 表单验证
func Validate(c *gin.Context, request FormRequest) bool {
	// 1. 解析请求，支持 JSON 数据、表单请求和 URL Query
	if err := c.ShouldBind(request); err != nil {
		response.Abort400(c, err)
		return false
	}
	// 2. 表单验证
	opts := govalidator.Options{
		Data:          request,
		Rules:         request.Rules(),
		Messages:      request.Messages(),
		TagIdentifier: "valid",
	}
	errs := govalidator.New(opts).ValidateStruct()

	// 3. 判断验证是否通过
	if len(errs) > 0 {
		response.Abort422(c, errs)
		return false
	}
	return true
}
