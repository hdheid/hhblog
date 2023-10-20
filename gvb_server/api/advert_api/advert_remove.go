package advert_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
)

// AdvertRemoveView 删除广告
// @tags 广告管理
// @Summary 删除广告
// @Description 通过id列表的形式一次性删除多个或者一个广告
// @Param data body models.RemoveRequest	true  "广告ID列表"
// @Router /api/adverts [delete]
// @Produce json
// @Success 200 {object} common.Response{}
func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Debug("参数解析失败：%s", err)
		common.FailWithCode(common.ArgumentError, c)
		return
	}

	var advertList []models.AdvertModel
	count := global.DB.Find(&advertList, cr.IDList).RowsAffected
	if count == 0 {
		common.FailWithMessage("广告不存在！", c)
		return
	}

	global.DB.Delete(&advertList) //删除数据库中的相应广告信息
	common.OKWithMessage(fmt.Sprintf("一共删除了%d个广告！", count), c)
}
