package main

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/core"
	"gvb_server/global"
)

var client *elastic.Client

func InitEs() {
	var err error
	sniffOpt := elastic.SetSniff(false)

	host := "http://127.0.0.1:9200"

	c, err := elastic.NewClient(
		elastic.SetURL(host),
		sniffOpt,
		elastic.SetBasicAuth("", ""),
	)
	if err != nil {
		global.Log.Fatalf("es连接失败：%s", err.Error())
	}

	client = c
}

func init() {
	core.InitConf()
	core.InitDefaultLogger()
	InitEs()
}

type DemoModel struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Key       string `json:"key"`
	UserID    uint   `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

func (DemoModel) Index() string {
	return "demo_index"
}

func Create(data *DemoModel) (err error) {
	indexResponse, err := client.Index().
		Index(data.Index()).
		BodyJson(data).
		Do(context.Background())

	if err != nil {
		logrus.Error("数据添加失败：%s", err.Error())
		return err
	}

	logrus.Infof("数据添加成功：%s", indexResponse.Id)
	data.ID = indexResponse.Id
	return nil
}

func FindList(key string, page, limit int) (demoList []DemoModel, count int) {
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

	res, err := client.Search(DemoModel{}.Index()).
		Query(boolSearch).
		From((from - 1) * limit).
		Size(limit).
		Do(context.Background())

	if err != nil {
		logrus.Error(err.Error())
		return
	}

	count = int(res.Hits.TotalHits.Value) //搜索结果的总条数
	for _, hit := range res.Hits.Hits {
		var demo DemoModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			logrus.Error(err.Error())
			return
		}
		err = json.Unmarshal(data, &demo)
		if err != nil {
			logrus.Error(err.Error())
			return
		}

		demo.ID = hit.Id
		demoList = append(demoList, demo)
	}

	return
}

func FindSourceList(key string, page, limit int) {

}

func Update(id string, data *DemoModel) error {
	_, err := client.Update().
		Index(DemoModel{}.Index()).
		Id(id).
		Doc(map[string]string{
			"title": data.Title,
		}).
		Do(context.Background())

	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	logrus.Infof("更新成功！")
	return nil
}

func Remove(idList []string) (count int, err error) {
	bulkService := client.Bulk().Index(DemoModel{}.Index()).Refresh("true")

	for _, id := range idList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)
	}

	res, err := bulkService.Do(context.Background())
	return len(res.Succeeded()), err
}

func main() {
	//Create(&DemoModel{
	//	Title:     "go开发",
	//	Key:       "",
	//	UserID:    2,
	//	CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	//})

	//lsit, count := FindList("", 1, 10)
	//fmt.Println(lsit, count)
	/*
		刚刚创建好的数据马上去查询是查询不到的，需要刷新一下
	*/

	//Update("O3wo8YsByY-YCAmsLwKA", &DemoModel{Title: "你好"})

	//Remove([]string{"PHwp8YsByY-YCAmsLwJZ"})
}
