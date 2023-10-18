package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
)

func (ImagesApi) ImageRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Debug("参数解析失败：%s", err)
		common.FailWithCode(common.ArgumentError, c)
		return
	}

	var imageList []models.BannerModel
	count := global.DB.Find(&imageList, cr.IDList).RowsAffected
	if count == 0 {
		common.FailWithMessage("图片不存在！", c)
		return
	}

	global.DB.Delete(&imageList) //删除数据库中的相应图片信息,绑定一个删除函数的钩子函数，将图片也删除
	common.OKWithMessage(fmt.Sprintf("一共删除了%d张图片！", count), c)
}
