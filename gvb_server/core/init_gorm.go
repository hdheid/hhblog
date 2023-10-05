package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/global"
	"time"
)

/*
使用gorm来连接数据
*/

func InitGorm() {
	//判断mysql有没有连接地址
	if global.Config.Mysql.Host == "" {
		global.Log.Warnln("数据库连接地址 host 是空的！")
	}

	//获取 dsn 连接 url
	dsn := global.Config.Mysql.Dsn()

	var mysqllogger logger.Interface
	//如果当前为开发环境，显示所有 SQL 语句
	if global.Config.System.Env == "dev" {
		mysqllogger = logger.Default.LogMode(logger.Info)
	} else {
		//只打印错误的sql语句
		mysqllogger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqllogger,
	})
	if err != nil {
		global.Log.Fatalf(fmt.Sprintf("[%s] [%s] mysql连接失败! ", dsn, err))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最多可容纳
	sqlDB.SetConnMaxLifetime(4 * time.Hour) // 连接最大复用时间，不能超过mysql的wait_timeout

	global.Log.Info("数据库连接成功!")
	global.DB = db
}
