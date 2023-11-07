package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
)

type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请输入图片id"`
	Name string `json:"name" binding:"required" msg:"请输入图片名称"`
}

func (ImagesApi) ImageUpdateView(c *gin.Context) {
	var cr ImageUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Debugf("参数解析失败：%s", err)
		common.FailWithError(err, &cr, c)
		return
	}

	//进行图片信息的修改
	var imageModel models.BannerModel
	err = global.DB.Take(&imageModel, cr.ID).Error //从数据库中获取图片的信息
	if err != nil {
		global.Log.Debug("数据查找失败！")
		common.FailWithMessage("该图片不存在！", c)
		return
	}

	err = global.DB.Model(&imageModel).Update("name", cr.Name).Error //进行数据库信息的修改
	if err != nil {
		global.Log.Debug("数据修改失败!")
		common.FailWithMessage(err.Error(), c)
		return
	}

	common.OKWithMessage("图片修改成功！", c)
}
