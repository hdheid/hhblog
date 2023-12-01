package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/es_ser"
)

func main() {
	core.InitConf()
	core.InitLogger()
	core.InitES()

	boolSearch := elastic.NewMatchAllQuery()
	res, _ := global.Client.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())

	for _, hit := range res.Hits.Hits {
		var article models.ArticleModel
		_ = json.Unmarshal(hit.Source, &article)
		article.ID = hit.Id

		indexList := es_ser.GetSearchIndexDataByContent(hit.Id, article.Title, article.Content)

		//批量添加
		bulkService := global.Client.Bulk()

		for _, data := range indexList {
			req := elastic.NewBulkIndexRequest().Index(models.FullTextModel{}.Index()).Doc(data)
			bulkService.Add(req)
		}

		result, err := bulkService.Do(context.Background())
		if err != nil {
			global.Log.Error(err)
		}

		fmt.Println("添加成功！", len(result.Succeeded()))
	}
}
