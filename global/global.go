package global

import (
	"github.com/go-redis/redis"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/config"
)

var (
	Config      *config.Config   //yaml读取的配置
	Log         *logrus.Logger   //日志的全局变量
	DB          *gorm.DB         //数据库的全局变量
	RDB         *redis.Client    //redis数据库的全局变量
	MysqlLogger logger.Interface //数据库日志的全局变量
	Client      *elastic.Client  //es的全局变量
)
