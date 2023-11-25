package article_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common"
	"gvb_server/service/es_ser"
)

type ESIdRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr ESIdRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		global.Log.Error(err.Error())
		common.FailWithCode(common.ArgumentError, c)
		return
	}

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
