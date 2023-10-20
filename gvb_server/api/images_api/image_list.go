package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/service/list_func"
)

/*
图片列表展示一般需要分页展示，这里会有分页展示功能
*/

func (ImagesApi) ImageListView(c *gin.Context) {
	var conf models.PageInfo
	err := c.ShouldBindQuery(&conf)
	if err != nil {
		global.Log.Debug("参数解析失败：%s", err)
		common.FailWithCode(common.ArgumentError, c)
		return
	}

	var banner models.BannerModel
	imageList, count, err := list_func.ComList(banner, list_func.Option{
		PageInfo: conf,
		Debug:    false,
	})

	common.OKWithList(imageList, count, c) //返回响应
}
