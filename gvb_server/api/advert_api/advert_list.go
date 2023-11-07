package advert_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/service/list_func"
	"strings"
)

// AdvertListView 广告的查询操作
// @tags 广告管理
// @Summary 查询广告
// @Description 可以分页查询广告列表
// @Param data query models.PageInfo	false  "查询参数"
// @Router /api/adverts [get]
// @Produce json
// @Success 200 {object} common.Response{data=common.ListResponse[models.AdvertModel]}
func (AdvertApi) AdvertListView(c *gin.Context) {
	var conf models.PageInfo
	err := c.ShouldBindQuery(&conf)
	if err != nil {
		global.Log.Debugf("参数解析失败：%s", err)
		common.FailWithCode(common.ArgumentError, c)
		return
	}

	//通过 referer 来判断是否包含admin，如果是，则展示所有，否则只展示 is_show 为 true 的部分
	referer := c.GetHeader("Referer")
	global.Log.Debug(referer)

	isShow := true
	if strings.Contains(referer, "admin") { //如果包含admin，就表示是 admin 过来地请求，就展示所有
		isShow = false
	}

	//当 is_show 为 true 的时候，只能查到 is_show 为 true 的数据，当is_show 为 false 的时候，不管是 true 还是 false 的数据都能够查得到
	advert := models.AdvertModel{
		IsShow: isShow,
	}
	advertList, count, err := list_func.ComList(advert, list_func.Option{
		PageInfo: conf,
		Debug:    false,
	})

	common.OKWithList(advertList, count, c) //返回响应
}
