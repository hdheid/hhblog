package advert_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
)

// AdvertUpdateView 广告的修改操作
// @tags 广告管理
// @Summary 修改广告
// @Description 查询ID，来修改该广告的内容
// @Param data body AdvertRequest	true  "广告的一些参数"
// @Router /api/adverts/:id [put]
// @Produce json
// @Success 200 {object} common.Response{}
func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	//参数绑定
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Debugf("参数解析失败：%s", err)
		common.FailWithError(err, &cr, c)
		return
	}

	//获取广告的 id
	id := c.Param("id")

	var advert models.AdvertModel
	err = global.DB.Take(&advert, id).Error
	if err != nil {
		common.FailWithMessage("广告不存在！", c)
		return
	}

	//global.Log.Warnln(cr)

	/*
		可以使用一个struct转map的第三方库：structs,很方便的实现struct对map的转化
		使用 go get github.com/fatih/structs
	*/

	//修改数据
	err = global.DB.Model(&models.AdvertModel{}).Where("id = ?", id).Updates(structs.Map(&cr)).Error

	if err != nil {
		global.Log.Error("更新广告失败：", err)
		common.FailWithMessage("更新广告失败！", c)
		return
	}

	common.OKWithMessage("更新广告成功！", c)

}
