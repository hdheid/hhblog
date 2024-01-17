package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
)

func (MenuApi) MenuDeleteView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Debugf("参数解析失败：%s", err)
		common.FailWithCode(common.ArgumentError, c)
		return
	}

	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, cr.IDList).RowsAffected
	if count == 0 {
		common.FailWithMessage("菜单不存在！", c)
		return
	}

	//使用事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		//首先删除关联的banners
		err = global.DB.Model(&menuList).Association("Banners").Clear()
		if err != nil {
			global.Log.Error("删除菜单关联图片失败！", err)
			return err
		}
		//删除菜单
		err = global.DB.Delete(&menuList).Error //删除数据库中的相应菜单信息
		if err != nil {
			global.Log.Error("删除菜单失败", err)
			return err
		}

		return nil
	})

	if err != nil {
		common.FailWithMessage("删除菜单失败！", c)
	}

	common.OKWithMessage(fmt.Sprintf("一共删除了%d个菜单！", count), c)
}
