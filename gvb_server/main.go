package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	core.InitConf() //读取配置文件
	fmt.Println(global.Config)
}
