package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"os"
	"path"
)

/*
使用的是 github.com/sirupsen/logrus 的日志打印
所以需要安装，使用 go get github.com/sirupsen/logrus
*/

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的级别来展示相应颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	log := global.Config.Logger

	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自义定文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自义定输出格式
		fmt.Fprintf(b, "[%s] [%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", log.Prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] [%s] \x1b[%dm[%s]\x1b[0m %s\n", log.Prefix, timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() {
	mLog := logrus.New()                                //新建一个实例
	mLog.SetOutput(os.Stdout)                           //设置输出类型
	mLog.SetReportCaller(global.Config.Logger.ShowLine) //开启是否返回函数名和行号
	mLog.SetFormatter(&LogFormatter{})                  //设置自定义的日志格式

	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil { //如果读取错误，设置为默认等级
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level) //设置日志级别，从配置文件中读取

	global.Log = mLog
	global.Log.Info("日志初始化成功！")
}

func InitDefaultLogger() {
	//全局log
	logrus.SetOutput(os.Stdout)                           //设置输出类型
	logrus.SetReportCaller(global.Config.Logger.ShowLine) //开启是否返回函数名和行号
	logrus.SetFormatter(&LogFormatter{})                  //设置自定义的日志格式

	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil { //如果读取错误，设置为默认等级
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level) //设置日志级别，从配置文件中读取
	logrus.Info("日志初始化成功！")
}
