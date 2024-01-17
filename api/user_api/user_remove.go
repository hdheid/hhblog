package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
)

func (UserApi) UserRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Debugf("参数解析失败：%s", err)
		common.FailWithCode(common.ArgumentError, c)
		return
	}

	var userList []models.UserModel
	count := global.DB.Find(&userList, cr.IDList).RowsAffected
	if count == 0 {
		common.FailWithMessage("用户不存在！", c)
		return
	}

	//使用事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		//删除菜单
		//TODO: 删除消息表、评论表、用户收藏的文章、发布的文章
		err = global.DB.Delete(&userList).Error //删除数据库中的相应菜单信息
		if err != nil {
			global.Log.Error("删除用户失败", err)
			return err
		}

		return nil
	})

	if err != nil {
		common.FailWithMessage("删除用户失败！", c)
	}

	common.OKWithMessage(fmt.Sprintf("一共删除了%d个用户！", count), c)
}
