package article_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/service/list_func"
	"gvb_server/utils/jwts"
)

type CollResponse struct {
	models.ArticleModel
	CreateAt string `json:"create_at"`
}

func (ArticleApi) ArticleCollListView(c *gin.Context) {
	client := global.Client

	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		global.Log.Error(err.Error())
		common.FailWithCode(common.ArgumentError, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims) //断言

	var articleIdList []interface{}
	var collTimeMap = map[string]string{}
	//var coll []models.UserCollectModel
	//global.DB.Select("article_id").Find(&coll, "user_id = ?", claims.UserID).Scan(&articleIdList)

	//传id列表，查询es
	list, count, _ := list_func.ComList(models.UserCollectModel{UserID: claims.UserID}, list_func.Option{
		PageInfo: cr,
	})
	for _, model := range list {
		articleIdList = append(articleIdList, model.ArticleID)
		collTimeMap[model.ArticleID] = model.CreatedAt.Format("2006-01-02 15:04:05")
	}

	boolSearch := elastic.NewTermsQuery("_id", articleIdList...)
	result, err := client.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err.Error())
		common.FailWithMessage(err.Error(), c)
	}

	var collList = make([]CollResponse, 0)
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err.Error())
			continue
		}

		article.ID = hit.Id

		collList = append(collList, CollResponse{
			ArticleModel: article,
			CreateAt:     collTimeMap[article.ID],
		})
	}

	common.OKWithList(collList, count, c)

}
