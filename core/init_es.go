package core

import (
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
)

func InitES() {
	var err error
	sniffOpt := elastic.SetSniff(false)
	c, err := elastic.NewClient(
		elastic.SetURL(global.Config.ES.URL()),
		sniffOpt,
		elastic.SetBasicAuth(global.Config.ES.User, global.Config.ES.Password),
	)
	if err != nil {
		global.Log.Fatalf("es连接失败：%s", err.Error())
	}

	global.Log.Info("ES 连接成功!")
	global.Client = c
}
