package main

import (
	"context"
	"github.com/sirupsen/logrus"
)

func (DemoModel) Mapping() string {
	return `
{
  "settings": {
    "index":{
      "max_result_window": "100000"
    }
  }, 
  "mappings": {
    "properties": {
      "title": { 
        "type": "text"
      },
      "key": { 
        "type": "keyword"
      },
      "user_id": {
        "type": "integer"
      },
      "created_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
`
}

func (demo DemoModel) IndexExists() bool {
	exists, err := client.IndexExists(demo.Index()).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return exists
	}
	return exists
}

func (demo DemoModel) CreateIndex() error {
	if demo.IndexExists() {
		//有索引，那就删掉索引，然后再创建
		demo.RemoveIndex()
	}
	//无索引，就可以直接创建索引
	createIndex, err := client.
		CreateIndex(demo.Index()).
		BodyString(demo.Mapping()).
		Do(context.Background())
	if err != nil {
		logrus.Error("创建索引失败：", err.Error())
		return err
	}
	if !createIndex.Acknowledged {
		logrus.Error("创建失败：", err.Error())
		return err
	}
	logrus.Infof("索引 %s 创建成功！", demo.Index())
	return nil
}

func (demo DemoModel) RemoveIndex() error {
	logrus.Info("索引存在，删除索引！")
	indexDelete, err := client.DeleteIndex(demo.Index()).Do(context.Background())
	if err != nil {
		logrus.Error("删除索引失败：", err.Error())
		return err
	}
	if !indexDelete.Acknowledged {
		logrus.Error("删除索引失败：", err.Error())
		return err
	}

	logrus.Info("删除索引成功！")
	return nil
}
