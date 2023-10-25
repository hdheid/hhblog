package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
)

func (MenuApi) MenuDetailView(c *gin.Context) {
	id := c.Param("id")
	var menu models.MenuModel
	err := global.DB.Take(&menu, id).Error
	if err != nil {
		global.Log.Debug("该菜单不存在！")
		common.FailWithMessage("该菜单不存在！", c)
		return
	}

	//查询第三张连接表
	var menuBanner []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanner, "menu_id = ?", id)

	//现在的目标是，通过第三张表，将menu对应的一排banner的id和路径存放到一起
	banners := make([]Banner, 0)
	for _, banner := range menuBanner {
		//找到和该menuId相关联的所有图片的数据
		if menu.ID != banner.MenuID {
			continue
		}

		banners = append(banners, Banner{
			BannerId: banner.BannerID,
			Path:     banner.BannerModel.Path,
		})
	}

	menuResponse := MenuResponse{
		Menu:   menu,
		Banner: banners,
	}

	common.OKWithData(menuResponse, c)
}
