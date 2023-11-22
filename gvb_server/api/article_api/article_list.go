package article_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/service/es_ser"
)

func (ArticleApi) ArticleListView(c *gin.Context) {
	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		global.Log.Error(err.Error())
		common.FailWithCode(common.ArgumentError, c)
		return
	}

	list, count, err := es_ser.CommonList(page.Key, page.Page, page.Limit)
	if err != nil {
		global.Log.Errorf("查询失败：%s", err.Error())
		common.FailWithMessage("查询失败！", c)
		return
	}

	common.OKWithList(filter.Omit("list", list), int64(count), c)
}
