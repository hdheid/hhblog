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

type CalenderResponse struct {
	Data  string `json:"data"`
	Count int    `json:"count"`
}

type BucketsType struct {
	Buckets []struct {
		KeyAsString string `json:"key_as_string"`
		Key         int64  `json:"key"`
		DocCount    int    `json:"doc_count"`
	} `json:"buckets"`
}

var DateCount = map[string]int{}

func (ArticleApi) ArticleCalendarView(c *gin.Context) {
	client := global.Client
	agg := elastic.NewDateHistogramAggregation().Field("created_at").CalendarInterval("day")

	//从今天开始到去年的今天的数据，时间段搜索
	now := time.Now()
	AYearAgo := now.AddDate(-1, 0, 0) //表示去年的现在
	query := elastic.NewRangeQuery("created_at").
		Gte(AYearAgo.Format(time.DateTime)).Lte(now.Format(time.DateTime)) //lt小于某个数，gt大于某个数，这个表示大于去年这个时间小于今年这个时间

	res, err := client.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("calendar", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error("查询失败！", err.Error())
		common.FailWithMessage("查询失败！", c)
		return
	}

	var data BucketsType
	err = json.Unmarshal(res.Aggregations["calendar"], &data)
	if err != nil {
		common.FailWithMessage("日历转换错误！", c)
		global.Log.Error("日历转换错误！", err.Error())
		return
	}

	var resList = make([]CalenderResponse, 0)
	for _, bucket := range data.Buckets {
		Time, _ := time.Parse(time.DateTime, bucket.KeyAsString)
		DateCount[Time.Format("2006-01-02")] = bucket.DocCount
	}

	days := int(now.Sub(AYearAgo).Hours() / 24)
	for i := 0; i <= days; i++ { //这里会少算一天，所以需要小于等于
		day := AYearAgo.AddDate(0, 0, i).Format("2006-01-02")
		count, _ := DateCount[day]
		resList = append(resList, CalenderResponse{
			Data:  day,
			Count: count,
		})
	}

	common.OKWithData(resList, c)
}
