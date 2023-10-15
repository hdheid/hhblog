package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/config"
)

var (
	Config      *config.Config   //yaml读取的配置
	Log         *logrus.Logger   //日志的全局变量
	DB          *gorm.DB         //数据库的全局变量
	MysqlLogger logger.Interface //数据库日志的全局变量
)
