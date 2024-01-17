package es_ser

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis_ser"
	"strings"
)

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
	diggInfo := redis_ser.NewDigg().GetInfo()
	lookInfo := redis_ser.NewArticleLook().GetInfo()

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
		a.ID = hit.Id

		//优化，如果数据没有变化，那么就不更新，节省性能
		if diggInfo[a.ID] == 0 {
			global.Log.Info("数据没有变化")
		} else {
			a.DiggCount = a.DiggCount + diggInfo[a.ID] //更新点赞量
		}

		//浏览量
		if lookInfo[a.ID] == 0 {
			global.Log.Info("数据没有变化")
		} else {
			a.LookCount = a.LookCount + lookInfo[a.ID]
		}

		articleList = append(articleList, a)
	}

	return articleList, count, nil
}

func CommonDetail(id string) (article models.ArticleModel, err error) {
	client := global.Client
	res, err := client.
		Get().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err.Error())
		return article, err
	}

	if res.Found == false { //补充逻辑，没找到文章不会报错
		return article, errors.New("文章不存在")
	}

	err = json.Unmarshal(res.Source, &article)
	if err != nil {
		global.Log.Error(err.Error())
		return article, err
	}

	article.ID = res.Id

	//文章浏览量
	look, _ := redis_ser.NewArticleLook().Get(article.ID)
	article.LookCount = article.LookCount + look

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

func ArticleUpdate(client *elastic.Client, id string, data map[string]any) error {
	_, err := client.
		Update().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Doc(data).
		Do(context.Background())

	return err
}
