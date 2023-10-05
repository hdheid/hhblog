package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/routers"
)

func main() {
	//读取配置文件
	core.InitConf()

	//初始化日志
	core.InitLogger()

	//初始化数据库
	core.InitGorm()

	r := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("程序运行在：%s", addr)
	r.Run(addr)
}
