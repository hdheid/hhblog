package es_ser

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"strings"
)

type Option struct {
	Page   int    `form:"page"`
	Key    string `form:"key"`
	Limit  int    `form:"limit"`
	Sort   string `form:"sort"`
	Fields []string
	Tag    string `form:"tag"`
}

type SortFiled struct {
	Field     string
	Ascending bool
}

func (o *Option) GetForm() int {
	if o.Page == 0 {
		o.Page = 1
	}
	if o.Limit == 0 {
		o.Limit = 10
	}

	return (o.Page - 1) * o.Limit
}

func CommonList(option Option) (articleList []models.ArticleModel, count int, err error) {
	client := global.Client

	boolSearch := elastic.NewBoolQuery()

	if option.Key != "" {
		boolSearch.Must(
			elastic.NewMultiMatchQuery(option.Key, option.Fields...),
		)
	}

	if option.Tag != "" {
		boolSearch.Must(
			elastic.NewMultiMatchQuery(option.Tag, "tags"),
		)
	}

	sortFiled := SortFiled{"created_at", false} //默认按照时间从近到远排列
	if option.Sort != "" {
		list := strings.Split(option.Sort, " ")
		if len(list) == 2 {
			sortFiled.Field = list[0]              //排序的key
			sortFiled.Ascending = list[1] == "asc" //从大到小还是从小到大
		}
	}

	res, err := client.Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		From(option.GetForm()).
		Highlight(elastic.NewHighlight().Field("title")). //es关于搜索高亮的一些东西，只搜索了标题
		Sort(sortFiled.Field, sortFiled.Ascending).
		Size(option.Limit).
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
		if title, ok := hit.Highlight["title"]; ok {
			a.Title = title[0]
		}
		fmt.Println(hit.Highlight)
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
