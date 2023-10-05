package main

import (
	"gvb_server/core"
)

func main() {
	core.InitConf() //读取配置文件
	//fmt.Println(global.Config)

	//初始化日志
	core.InitLogger()

	//初始化数据库
	core.InitGorm()

}
