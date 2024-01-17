package article_api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/service/es_ser"
	"gvb_server/utils/jwts"
)

func (ArticleApi) ArticleCollBatchRemoveView(c *gin.Context) {
	client := global.Client
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims) //断言

	var cr models.ESIdListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		common.FailWithCode(common.ArgumentError, c)
		return
	}

	//查询出来
	var collects []models.UserCollectModel
	var articleIdList []string
	global.DB.Find(&collects, "user_id = ? and article_id in ?", claims.UserID, cr.IDList).Select("article_id").Scan(&articleIdList)
	if len(articleIdList) == 0 {
		global.Log.Error("没有收藏的文章")
		common.OKWithMessage("没有收藏的文章！", c)
		return
	}

	var idList []interface{}
	for _, i := range articleIdList {
		idList = append(idList, i)
	}

	//更新文章数量
	boolSearch := elastic.NewTermsQuery("_id", idList...)
	result, err := client.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err.Error())
		common.FailWithMessage(err.Error(), c)
	}

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err.Error())
			continue
		}
		err = es_ser.ArticleUpdate(client, hit.Id, map[string]any{
			"collects_count": article.CollectsCount - 1,
		})
		if err != nil {
			global.Log.Error(err.Error())
			continue
		}
	}

	global.DB.Delete(&collects)

	common.OKWithMessage(fmt.Sprintf("成功取消收藏 %d 篇文章", len(articleIdList)), c)
}
