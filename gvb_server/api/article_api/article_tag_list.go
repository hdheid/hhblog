package article_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"time"
)

type TagResponse struct {
	Tag           string   `json:"tag"`
	Count         int      `json:"count"`
	ArticleIDList []string `json:"article_id_list"`
	CreatedAt     string   `json:"created_at"`
}

type TagsType struct {
	DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int `json:"sum_other_doc_count"`
	Buckets                 []struct {
		Key      string `json:"key"`
		DocCount int    `json:"doc_count"`
		Articles struct {
			DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int `json:"sum_other_doc_count"`
			Buckets                 []struct {
				Key      string `json:"key"`
				DocCount int    `json:"doc_count"`
			} `json:"buckets"`
		} `json:"articles"`
	} `json:"buckets"`
}

func (ArticleApi) ArticleTagListView(c *gin.Context) {
	var conf models.PageInfo
	err := c.ShouldBindQuery(&conf)
	if err != nil {
		global.Log.Debugf("参数解析失败：%s", err)
		common.FailWithCode(common.ArgumentError, c)
		return
	}

	if conf.Limit == 0 {
		conf.Limit = 10
	}
	offset := (conf.Page - 1) * conf.Limit
	if offset < 0 {
		offset = 0
	}

	client := global.Client

	res, err := client.
		Search(models.ArticleModel{}.Index()).
		Aggregation("tags", elastic.NewCardinalityAggregation().Field("tags")). //避免重复值
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error("查询失败！", err.Error())
		common.FailWithMessage("查询失败！", c)
		return
	}
	cTag, _ := res.Aggregations.Cardinality("tags")
	count := int64(*cTag.Value)

	agg := elastic.NewTermsAggregation().Field("tags")
	agg.SubAggregation("articles", elastic.NewTermsAggregation().Field("keyword"))
	agg.SubAggregation("pages", elastic.NewBucketSortAggregation().From(offset).Size(conf.Limit))

	query := elastic.NewBoolQuery()

	res, err = client.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error("查询失败！", err.Error())
		common.FailWithMessage("查询失败！", c)
		return
	}
	var tagType TagsType
	var tagList = make([]*TagResponse, 0)
	var tagStringList []string
	err = json.Unmarshal(res.Aggregations["tags"], &tagType)
	if err != nil {
		common.FailWithMessage("错误！", c)
		global.Log.Error("错误！", err.Error())
		return
	}

	for _, bucket := range tagType.Buckets {

		var articleList []string
		for _, article := range bucket.Articles.Buckets {
			articleList = append(articleList, article.Key)
		}

		tagList = append(tagList, &TagResponse{
			Tag:           bucket.Key,
			Count:         bucket.DocCount,
			ArticleIDList: articleList,
		})
		tagStringList = append(tagStringList, bucket.Key)
	}

	var tagModelList []models.TagModel
	err = global.DB.Find(&tagModelList, "title in ?", tagStringList).Error
	if err != nil {
		common.FailWithMessage("查询失败！", c)
		global.Log.Error("查询失败：%s", err.Error())
		return
	}

	//把标签创建时间搜索出来放到 tagList 中
	var tagDate = map[string]string{}
	for _, tag := range tagModelList {
		tagDate[tag.Title] = tag.CreatedAt.Format(time.DateTime)
	}
	for _, tag := range tagList {
		tag.CreatedAt = tagDate[tag.Tag]
	}

	common.OKWithList(tagList, count, c)
}
