package main

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis_ser"
)

func main() {
	core.InitConf()
	core.InitLogger()
	core.InitRedis()
	core.InitES()

	res, err := global.Client.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000). //默认情况下，size大小为10，Elasticsearch 将返回最多 10000 条匹配的文档作为搜索结果
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		return
	}

	diggInfo := redis_ser.NewDigg().GetInfo()        //获取数据
	lookInfo := redis_ser.NewArticleLook().GetInfo() //获取浏览量
	for _, hit := range res.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article) //这里article的ID并没有被赋值
		if err != nil {
			global.Log.Error(err)
			continue
		}

		article.DiggCount = article.DiggCount + diggInfo[hit.Id] //更新点赞量
		article.LookCount = article.LookCount + lookInfo[hit.Id] //更新浏览量
		if lookInfo[hit.Id] == 0 && diggInfo[hit.Id] == 0 {
			global.Log.Info("数据没有变化")
			continue
		} //优化，如果数据没有变化，那么就不更新，节省性能

		_, err := global.Client. //更新es部分
						Update().
						Index(models.ArticleModel{}.Index()).
						Id(hit.Id).
						Doc(map[string]int{
				"digg_count": article.DiggCount,
				"look_count": article.LookCount,
			}).Do(context.Background())
		if err != nil {
			global.Log.Error(err)
			continue
		}
		global.Log.Infof("%s 的点赞量和浏览数同步成功！点赞数为：%d，浏览数为：%d", article.Title, article.DiggCount, article.LookCount)
	}

	redis_ser.NewDigg().Clear()        //删除redis的点赞数据
	redis_ser.NewArticleLook().Clear() //删除redis的浏览数据
}

/*
为什么对缓存只删除不更新？

不更新缓存是防止并发更新导致的数据不一致。
所以为了降低数据不一致的概率，不应该更新缓存，而是直接将其删除，
然后等待下次发生cache miss时再把数据库中的数据同步到缓存。

https://www.zhihu.com/question/23401553?sort=created
*/
