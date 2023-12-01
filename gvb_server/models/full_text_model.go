package models

import (
	"context"
	"gvb_server/global"
)

type FullTextModel struct {
	ID    string `json:"id" structs:"id"`       // es的id
	Key   string `json:"key" structs:"key"`     // 文章关联的id
	Title string `json:"title" structs:"title"` // 文章标题
	Slug  string `json:"slug" struct:"slug"`
	Body  string `json:"body" structs:"body"` // 文章内容
}

func (FullTextModel) Index() string {
	return "full_text_index"
}

func (FullTextModel) Mapping() string {
	return `
{
  "settings": {
    "index":{
      "max_result_window": "100000"
    }
  }, 
  "mappings": {
    "properties": {
      "key": { 
        "type": "keyword"
      },
      "title": { 
        "type": "text"
      },
      "slug": { 
        "type": "keyword"
      },
      "body": { 
        "type": "text"
      }
    }
  }
}
`
}

func (a FullTextModel) IndexExists() bool {
	client := global.Client
	exists, err := client.IndexExists(a.Index()).Do(context.Background())
	if err != nil {
		global.Log.Error(err.Error())
		return exists
	}
	return exists
}

func (a FullTextModel) CreateIndex() error {
	client := global.Client
	if a.IndexExists() {
		//有索引，那就删掉索引，然后再创建
		a.RemoveIndex()
	}
	//无索引，就可以直接创建索引
	createIndex, err := client.
		CreateIndex(a.Index()).
		BodyString(a.Mapping()).
		Do(context.Background())
	if err != nil {
		global.Log.Error("创建索引失败：", err.Error())
		return err
	}
	if !createIndex.Acknowledged {
		global.Log.Error("创建失败：", err.Error())
		return err
	}
	global.Log.Infof("索引 %s 创建成功！", a.Index())
	return nil
}

func (a FullTextModel) RemoveIndex() error {
	client := global.Client
	global.Log.Info("索引存在，删除索引！")
	indexDelete, err := client.DeleteIndex(a.Index()).Do(context.Background())
	if err != nil {
		global.Log.Error("删除索引失败：", err.Error())
		return err
	}
	if !indexDelete.Acknowledged {
		global.Log.Error("删除索引失败：", err.Error())
		return err
	}

	global.Log.Info("删除索引成功！")
	return nil
}
