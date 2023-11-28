package models

import (
	"context"
	"errors"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models/ctype"
)

type ArticleModel struct {
	ID        string `json:"id" structs:"id"`                 // es的id
	CreatedAt string `json:"created_at" structs:"created_at"` // 创建时间
	UpdatedAt string `json:"updated_at" structs:"updated_at"` // 更新时间

	Title    string `json:"title" structs:"title"`                // 文章标题
	Keyword  string `json:"keyword,omit(list)" structs:"keyword"` // 关键字
	Abstract string `json:"abstract" structs:"abstract"`          // 文章简介
	Content  string `json:"content,omit(list)" structs:"content"` // 文章内容

	LookCount     int `json:"look_count" structs:"look_count"`         // 浏览量
	CommentCount  int `json:"comment_count" structs:"comment_count"`   // 评论量
	DiggCount     int `json:"digg_count" structs:"digg_count"`         // 点赞量
	CollectsCount int `json:"collects_count" structs:"collects_count"` // 收藏量

	UserID       uint   `json:"user_id" structs:"user_id"`               // 用户id
	UserNickName string `json:"user_nick_name" structs:"user_nick_name"` //用户昵称
	UserAvatar   string `json:"user_avatar" structs:"user_avatar"`       // 用户头像

	Category string `json:"category" structs:"category"` // 文章分类
	Source   string `json:"source" structs:"source"`     // 文章来源
	Link     string `json:"link" structs:"link"`         // 原文链接

	BannerID  uint   `json:"banner_id" structs:"banner_id"`   // 文章封面id
	BannerUrl string `json:"banner_url" structs:"banner_url"` // 文章封面

	Tags ctype.Array `json:"tags" structs:"tags"` // 文章标签
}

func (ArticleModel) Index() string {
	return "article_index"
}

func (ArticleModel) Mapping() string {
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
      "keyword": { 
        "type": "keyword"
      },
      "abstract": { 
        "type": "text"
      },
      "content": { 
        "type": "text"
      },
      "look_count": {
        "type": "integer"
      },
      "comment_count": {
        "type": "integer"
      },
      "digg_count": {
        "type": "integer"
      },
      "collects_count": {
        "type": "integer"
      },
      "user_id": {
        "type": "integer"
      },
      "user_nick_name": { 
        "type": "keyword"
      },
      "user_avatar": { 
        "type": "keyword"
      },
      "category": { 
        "type": "keyword"
      },
      "source": { 
        "type": "keyword"
      },
      "link": { 
        "type": "keyword"
      },
      "banner_id": {
        "type": "integer"
      },
      "banner_url": { 
        "type": "keyword"
      },
      "tags": { 
        "type": "keyword"
      },
      "created_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      },
      "updated_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
`
}

func (a ArticleModel) IndexExists() bool {
	client := global.Client
	exists, err := client.IndexExists(a.Index()).Do(context.Background())
	if err != nil {
		global.Log.Error(err.Error())
		return exists
	}
	return exists
}

func (a ArticleModel) CreateIndex() error {
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

func (a ArticleModel) RemoveIndex() error {
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

//es文章的操作

func (a *ArticleModel) Create() (err error) {
	client := global.Client
	indexResponse, err := client.Index().
		Index(a.Index()).
		BodyJson(a).
		Refresh("true").
		Do(context.Background())
	if err != nil {
		global.Log.Error(err.Error())
		return err
	}
	a.ID = indexResponse.Id
	return nil
}

// ISExistData 是否存在该文章
func (a ArticleModel) ISExistData() bool {
	client := global.Client
	res, err := client.
		Search(a.Index()).
		Query(elastic.NewTermQuery("keyword", a.Title)).
		Size(1).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err.Error())
		return false
	}
	if res.Hits.TotalHits.Value > 0 {
		return true
	}
	return false
}

func (a *ArticleModel) GetDataByID(id string) error {
	client := global.Client
	res, err := client.
		Get().
		Index(a.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		return err
	}
	if res.Found == false { //补充逻辑，没找到文章不会报错
		return errors.New("文章不存在")
	}

	//err = json.Unmarshal(res.Source, a) //这里将数据传到a中后，是否就会出现问题？暂时先不需要这个
	return nil
}
