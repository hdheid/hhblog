package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
)

func main() {
	//读取配置文件
	core.InitConf()

	//初始化日志
	core.InitLogger()

	//初始化数据库
	core.InitGorm()

	var banner models.BannerModel
	err := global.DB.Take("id = ?", 1000).Error
	if err != nil {
		fmt.Println("图片不存在！")
	}

	fmt.Println(banner)
}
