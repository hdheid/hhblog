package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/models/ctype"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"请完善菜单名称" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"请完善菜单路径" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`           // 简介
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"` // 切换的时间，单位秒
	BannerTime    int         `json:"banner_time" structs:"banner_time"`     // 切换的时间，单位秒
	Sort          int         `json:"sort" structs:"sort"`                   // 菜单的序号
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`           // 具体图片的顺序
}

// MenuCreateView 菜单的添加操作
// @tags 菜单管理
// @Summary 创建菜单
// @Description 创建一个或多个菜单
// @Param data body MenuRequest	  true  "表示多个参数"
// @Router /api/menus [post]
// @Produce json
// @Success 200 {object} common.Response{}
func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Debugf("参数解析失败:%s", err)
		common.FailWithError(err, &cr, c)
	}

	//入库前判重
	var menu models.MenuModel
	err = global.DB.Take(&menu, "title = ? or path = ?", cr.Title, cr.Path).Error
	if err == nil {
		common.FailWithMessage("菜单已存在，请勿重复添加！", c)
		return
	}

	//创建banner数据入库
	menuModel := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}

	err = global.DB.Create(&menuModel).Error
	if err != nil {
		global.Log.Error("创建菜单失败：", err)
		common.FailWithMessage("添加菜单失败！", c)
		return
	}

	//第三张表数据入库
	if len(cr.ImageSortList) == 0 { //如果没有图片，就没有数据入库第三张表
		common.OKWithMessage("添加菜单成功！", c)
		return
	}

	var menuBannerList []models.MenuBannerModel

	for _, menuBanner := range cr.ImageSortList {
		//后期需要判断image_id是否存在该图片
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: menuBanner.ImageID,
			Sort:     menuBanner.Sort,
		})
	}

	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		global.Log.Error(err)
		common.FailWithMessage("添加菜单图片失败", c)
		return
	}

	common.OKWithMessage("添加菜单成功！", c)
}
