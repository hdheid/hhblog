package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 常量封装
const (
	SUCCESS = 0
	ERR     = 1
)

// 封装响应
type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

// 将响应封装并返回
func result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

// OK 响应成功
func OK(data any, msg string, c *gin.Context) {
	result(SUCCESS, data, msg, c)
}

// OKWithNoting 什么都不需要传
func OKWithNoting(c *gin.Context) {
	result(SUCCESS, map[string]interface{}{}, "成功", c)
}

// OKWithMessage 只需要传msg
func OKWithMessage(msg string, c *gin.Context) {
	result(SUCCESS, map[string]interface{}{}, msg, c)
}

// OKWithData 只需要传data
func OKWithData(data any, c *gin.Context) {
	result(SUCCESS, data, "成功", c)
}

// Fail 失败
func Fail(data any, msg string, c *gin.Context) {
	result(SUCCESS, data, msg, c)
}

// FailWithMessage 只需要传msg
func FailWithMessage(msg string, c *gin.Context) {
	result(ERR, map[string]interface{}{}, msg, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	if msg, ok := ErrMap[code]; ok { //如果错误码能够查询到就响应回去
		result(int(code), map[string]interface{}{}, msg, c)
		return
	}
	result(ERR, map[string]interface{}{}, "未知错误", c) //如果是map中没有的，那就是未知错误
}
