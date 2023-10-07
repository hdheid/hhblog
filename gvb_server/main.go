package main

import (
	"gvb_server/core"
	"gvb_server/flag"
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

	//命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	r := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("程序运行在：%s", addr)
	err := r.Run(addr) //防止出错，使得程序更加健壮
	if err != nil {
		global.Log.Fatalf("程序启动失败：%s", err)
	}
}
