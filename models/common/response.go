package common

import (
	"github.com/gin-gonic/gin"
	"gvb_server/utils"
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

type ListResponse[T any] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
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

func OKWithList(list any, count int64, c *gin.Context) {

	result(SUCCESS, ListResponse[any]{
		List:  list,
		Count: count,
	}, "成功", c)
}

// Fail 失败
func Fail(data any, msg string, c *gin.Context) {
	result(SUCCESS, data, msg, c)
}

// FailWithMessage 只需要传msg
func FailWithMessage(msg string, c *gin.Context) {
	result(ERR, map[string]interface{}{}, msg, c)
}

func FailWithError(err error, obj any, c *gin.Context) {
	msg := utils.GetValidMsg(err, obj) //通过 GetValidMsg 函数获取到结构体后面的报错信息，是字符串形式的
	FailWithMessage(msg, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	if msg, ok := ErrMap[code]; ok { //如果错误码能够查询到就响应回去
		result(int(code), map[string]interface{}{}, msg, c)
		return
	}
	result(ERR, map[string]interface{}{}, "未知错误", c) //如果是map中没有的，那就是未知错误
}
