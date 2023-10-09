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

	//可以像这样路由分组
	apiRouterGroup := r.Group("api")
	{
		SettingsRouter(apiRouterGroup) //获取系统信息接口
		ImagesRouter(apiRouterGroup)   //上传图片接口
	}

	return r
}
