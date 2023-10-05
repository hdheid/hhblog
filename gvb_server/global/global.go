package global

import (
	"github.com/sirupsen/logrus"
	"gvb_server/config"
)

var (
	Config *config.Config //yaml读取的配置
	Log    *logrus.Logger //日志的全局变量
)
