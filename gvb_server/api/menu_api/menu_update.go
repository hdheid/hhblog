package menu_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
)

// MenuUpdateView 菜单的修改操作
// @tags 菜单管理
// @Summary 修改菜单
// @Description 查询ID，来修改该菜单的内容
// @Param data body MenuRequest	true  "菜单的一些参数"
// @Router /api/menu/:id [put]
// @Produce json
// @Success 200 {object} common.Response{}
func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Debug("参数解析失败:%s", err)
		common.FailWithError(err, &cr, c)
	}
	id := c.Param("id")
	//更新的逻辑为：先清空banner，如果选择了就重新添加，没有选择就等于清空了

	//使用级联删除来清空这个menu的所有的banner,清空之前先判断该菜单是否存在
	var menu models.MenuModel
	err = global.DB.Take(&menu, id).Error
	if err != nil {
		global.Log.Debug("该菜单不存在")
		common.FailWithMessage("该菜单不存在！", c)
		return
	}
	//清空banner
	err = global.DB.Model(&menu).Association("Banners").Clear()
	if err != nil {
		global.Log.Debug("清空banner失败!")
	}

	//判断修改后的菜单是否又选择了图片
	if len(cr.ImageSortList) > 0 {
		var bannerList []models.MenuBannerModel
		for _, sort := range cr.ImageSortList {
			bannerList = append(bannerList, models.MenuBannerModel{
				MenuID:   menu.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}

		err = global.DB.Create(&bannerList).Error
		if err != nil {
			global.Log.Error(err)
			common.FailWithMessage("修改菜单图片失败！", c)
			return
		}
	}

	mp := structs.Map(&cr)
	err = global.DB.Model(&menu).Updates(mp).Error
	if err != nil {
		global.Log.Error(err)
		common.FailWithMessage("菜单修改失败！", c)
		return
	}
	common.OKWithMessage("菜单修改成功！", c)
}
