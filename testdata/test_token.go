package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/utils/jwts"
)

func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	core.InitLogger()

	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: "我是大米",
		Role:     1,
		UserID:   1,
		Avatar:   "/static/avatar/default.jpg",
	})
	global.Log.Infoln(token)

	if err == nil {
		claims, err := jwts.ParseToken(token)
		global.Log.Infoln(claims, err)
	}
}
