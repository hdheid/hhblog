package tag_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
)

type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"请输入标题" structs:"title"` // 显示的标题
}

func (TagApi) TagCreateView(c *gin.Context) {
	//参数绑定
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Debugf("参数解析失败：%s", err)
		common.FailWithError(err, &cr, c)
		return
	}

	var tag models.TagModel
	err = global.DB.Take(&tag, "title = ?", cr.Title).Error
	if err == nil {
		common.FailWithMessage("标签已存在，请勿重复添加！", c)
		return
	}

	//添加进入数据库
	err = global.DB.Create(&models.TagModel{
		Title: cr.Title,
	}).Error

	if err != nil {
		global.Log.Error("创建标签失败：", err)
		common.FailWithMessage("添加标签失败！", c)
		return
	}

	common.OKWithMessage("添加标签成功！", c)

}
