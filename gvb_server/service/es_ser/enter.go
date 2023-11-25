package es_ser

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
)

func CommonList(key string, page, limit int) (articleList []models.ArticleModel, count int, err error) {
	client := global.Client

	boolSearch := elastic.NewBoolQuery()
	from := page
	if key != "" {
		boolSearch.Must(
			elastic.NewMatchQuery("title", key),
		)
	}

	if limit == 0 {
		limit = 10
	}
	if from == 0 {
		from = 1
	}

	res, err := client.Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		From((from - 1) * limit).
		Size(limit).
		Do(context.Background())

	if err != nil {
		global.Log.Error(err.Error())
		return articleList, count, err
	}

	count = int(res.Hits.TotalHits.Value) //搜索结果的总条数
	for _, hit := range res.Hits.Hits {
		var a models.ArticleModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			global.Log.Error(err.Error())
			return articleList, count, err
		}
		err = json.Unmarshal(data, &a)
		if err != nil {
			global.Log.Error(err.Error())
			return articleList, count, err
		}

		a.ID = hit.Id
		articleList = append(articleList, a)
	}

	return articleList, count, nil
}

func CommonDetail(id string) (article models.ArticleModel, err error) {
	client := global.Client
	res, err := client.Get().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err.Error())
		return article, err
	}

	err = json.Unmarshal(res.Source, &article)
	if err != nil {
		global.Log.Error(err.Error())
		return article, err
	}

	article.ID = res.Id

	return article, nil
}

func CommonDetailByKeyword(key string) (article models.ArticleModel, err error) {
	client := global.Client
	res, err := client.Search().
		Index(models.ArticleModel{}.Index()).
		Query(elastic.NewTermQuery("keyword", key)).
		Size(1).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err.Error())
		return article, err
	}

	if res.Hits.TotalHits.Value == 0 {
		return article, errors.New("文章不存在！")
	}

	hit := res.Hits.Hits[0]

	err = json.Unmarshal(hit.Source, &article)
	if err != nil {
		global.Log.Error(err.Error())
		return article, err
	}

	article.ID = hit.Id

	return article, nil
}
