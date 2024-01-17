package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/service/es_ser"
	"gvb_server/utils/jwts"
)

func (ArticleApi) ArticleCollCreateView(c *gin.Context) {
	client := global.Client
	var cr models.ESIdRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		common.FailWithCode(common.ArgumentError, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims) //断言

	article, err := es_ser.CommonDetail(cr.ID)
	if err != nil {
		global.Log.Error(err.Error())
		common.FailWithMessage("文章不存在！", c)
		return
	}

	var coll models.UserCollectModel
	err = global.DB.Take(&coll, "user_id = ? and article_id = ?", claims.UserID, cr.ID).Error
	if err != nil {
		//表示没有找到这个数据，那么就增加
		global.DB.Create(&models.UserCollectModel{
			UserID:    claims.UserID,
			ArticleID: cr.ID,
		})

		//文章收藏数+1
		err = es_ser.ArticleUpdate(client, cr.ID, map[string]any{
			"collects_count": article.CollectsCount + 1,
		})
		common.OKWithMessage("文章收藏成功！", c)
		return
	}

	//文章数收藏数-1
	global.DB.Delete(&coll)

	//更新es
	err = es_ser.ArticleUpdate(client, cr.ID, map[string]any{
		"collects_count": article.CollectsCount - 1,
	})

	common.OKWithMessage("取消收藏成功！", c)
}
