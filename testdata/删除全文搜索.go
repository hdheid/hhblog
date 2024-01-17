package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
)

func main() {
	core.InitConf()
	core.InitLogger()
	core.InitES()

	boolSearch := elastic.NewTermQuery("key", "1HkyJIwBphBkeiGYH7gB")

	res, _ := global.Client.
		DeleteByQuery().
		Index(models.FullTextModel{}.Index()).
		Query(boolSearch).
		Do(context.Background())
	fmt.Println(res.Deleted)
}
