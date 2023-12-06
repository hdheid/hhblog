package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/service/es_ser"
	"gvb_server/service/redis_ser"
)

func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr models.ESIdRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		global.Log.Error(err.Error())
		common.FailWithCode(common.ArgumentError, c)
		return
	}

	redis_ser.NewArticleLook().Set(cr.ID) //当调用这个接口的时候，就表示该文章被浏览了一次，因此浏览量可以加一

	article, err := es_ser.CommonDetail(cr.ID)
	if err != nil {
		global.Log.Errorf("查询失败：%s", err.Error())
		common.FailWithMessage("文章不存在！", c)
		return
	}

	common.OKWithData(article, c)
}

type ArticleDetailRequest struct {
	Title string `json:"title" form:"title"`
}

func (ArticleApi) ArticleDetaiByTitlelView(c *gin.Context) {
	var cr ArticleDetailRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		global.Log.Error(err.Error())
		common.FailWithCode(common.ArgumentError, c)
		return
	}

	article, err := es_ser.CommonDetailByKeyword(cr.Title)
	if err != nil {
		global.Log.Errorf("查询失败：%s", err.Error())
		common.FailWithMessage("文章不存在！", c)
		return
	}

	common.OKWithData(article, c)
}
