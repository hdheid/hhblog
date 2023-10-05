package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
)

/*
使用gin框架,需要获取依赖
go get github.com/gin-gonic/gin
*/

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env) //设置gin环境

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	return r
}
