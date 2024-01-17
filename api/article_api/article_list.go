package article_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/service/es_ser"
)

type ArticleSearchRequest struct {
	models.PageInfo
	Tag string `json:"tag" form:"tag"`
}

func (ArticleApi) ArticleListView(c *gin.Context) {
	var page ArticleSearchRequest
	if err := c.ShouldBindQuery(&page); err != nil {
		global.Log.Error(err.Error())
		common.FailWithCode(common.ArgumentError, c)
		return
	}

	list, count, err := es_ser.CommonList(es_ser.Option{
		Page:   page.Page,
		Key:    page.Key,
		Limit:  page.Limit,
		Sort:   page.Sort,
		Fields: []string{"title", "content"},
		Tag:    page.Tag,
	})
	if err != nil {
		global.Log.Errorf("查询失败：%s", err.Error())
		common.FailWithMessage("查询失败！", c)
		return
	}

	//避免空值问题
	dataList := filter.Omit("list", list)
	_list, _ := dataList.(filter.Filter)
	__list, _ := _list.MarshalJSON()
	if string(__list) == "{}" {
		dataList = make([]models.ArticleModel, 0)
	}

	common.OKWithList(dataList, int64(count), c)
}
