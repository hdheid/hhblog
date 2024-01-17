package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
)

//这一部分的数据整理有待优化

type Banner struct {
	BannerId uint   `json:"banner_id"`
	Path     string `json:"path"`
}

type MenuResponse struct {
	Menu   models.MenuModel `json:"menu_list"`
	Banner []Banner         `json:"banner"`
}

func (MenuApi) MenuListView(c *gin.Context) {
	//查询所有菜单，并将id单独放在一个切片中
	var menuList []models.MenuModel
	var menuIdList []uint
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIdList)

	//查询第三张连接表
	var menuBanner []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanner, "menu_id in ?", menuIdList)

	//现在的目标是，通过第三张表，将每一个menu对应的一排banner的id和路径存放到一起
	var menus []MenuResponse
	for _, menu := range menuList {
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

		menus = append(menus, MenuResponse{
			Menu:   menu,
			Banner: banners,
		})
	}

	common.OKWithData(menus, c)
}
