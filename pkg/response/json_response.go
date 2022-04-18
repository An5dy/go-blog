package response

import (
	"go-blog/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ResponseFormatterParams 响应数据格式
type ResponseFormatterParams struct {
	Code    int                 `json:"code,omitempty"`    // 业务响应码
	Data    interface{}         `json:"data,omitempty"`    // 返回数据
	Error   string              `json:"error,omitempty"`   // 错误消息
	Errors  map[string][]string `json:"errors,omitempty"`  // 错误信息
	Message string              `json:"message,omitempty"` // 返回信息
	Success bool                `json:"success"`           // 操作是否成功
}

// JSON 200 响应 JSON 数据
func JSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// Succeed 响应成功
func Succeed(c *gin.Context) {
	JSON(c, ResponseFormatterParams{
		Success: true,
		Message: "操作成功。",
	})
}

// Data 200 响应 data 数据
func Data(c *gin.Context, data interface{}) {
	JSON(c, ResponseFormatterParams{
		Success: true,
		Data:    data,
	})
}

// Created 201 响应创建数据
func Created(c *gin.Context, location ...string) {
	if len(location) > 0 {
		c.Header("Location", location[0])
	}
	c.Status(http.StatusCreated)
}

// NoContent 204 响应
func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// Abort 终止响应
func Abort(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, ResponseFormatterParams{
		Message: message,
		Success: false,
	})
}

// Abort400 请求解析错误
func Abort400(c *gin.Context, err error, message ...string) {
	logger.LogIf(err)
	c.AbortWithStatusJSON(http.StatusBadRequest, ResponseFormatterParams{
		Message: defaultMessage("请求解析错误。", message...),
	})
}

// Abort401 401 用户未认证
func Abort401(c *gin.Context, message ...string) {
	Abort(c, http.StatusUnauthorized, defaultMessage("未登录。", message...))
}

// Abort403 403 请求的权限不足
func Abort403(c *gin.Context, message ...string) {
	Abort(c, http.StatusForbidden, defaultMessage("请求权限不足。", message...))
}

// Abort404 404 请求数据不存在
func Abort404(c *gin.Context, message ...string) {
	Abort(c, http.StatusNotFound, defaultMessage("请求数据不存在。", message...))
}

// Abort405 405 请求方法不存在
func Abort405(c *gin.Context, message ...string) {
	Abort(c, http.StatusMethodNotAllowed, defaultMessage("请求方法不存在。", message...))
}

// Abort422 422 请求格式不正确
func Abort422(c *gin.Context, errors map[string][]string, message ...string) {
	c.JSON(http.StatusUnprocessableEntity, ResponseFormatterParams{
		Message: defaultMessage("请求格式不正确。", message...),
		Errors:  errors,
	})
}

// Abort500 服务器内部错误
func Abort500(c *gin.Context, message ...string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, ResponseFormatterParams{
		Success: false,
		Message: defaultMessage("服务器内部错误，请稍后再试。", message...),
	})
}

// Error 响应 404 或 422
func Error(c *gin.Context, err error, message ...string) {
	switch err {
	case gorm.ErrRecordNotFound:
		Abort404(c)
		return
	default:
		logger.LogIf(err)
		c.JSON(http.StatusUnprocessableEntity, ResponseFormatterParams{
			Message: defaultMessage("请求处理失败，请查看 error 的值。", message...),
			Error:   err.Error(),
		})
	}
}

// defaultMessage 默认消息
func defaultMessage(defaultMessage string, message ...string) string {
	if len(message) > 0 {
		return message[0]
	} else {
		return defaultMessage
	}
}
