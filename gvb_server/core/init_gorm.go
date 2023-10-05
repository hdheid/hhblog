package core

import (
	"gvb_server/global"
	"log"
)

/*
使用gorm来连接数据
*/

func InitGorm() {
	//判断mysql有没有连接地址
	if global.Config.Mysql.Host == "" {
		log.Panicln("数据库连接地址 host 是空的！")
	}
	//获取 dsn 连接 url
	//dsn := global.Config.Mysql.Dsn()

	//var mysqllogger log.Interface
	//如果当前为开发环境，显示所有 SQL 语句
	if global.Config.System.Env == "dev" {

	}
}
