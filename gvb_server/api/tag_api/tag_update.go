package tag_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
)

func (TagApi) TagUpdateView(c *gin.Context) {
	//参数绑定
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Debugf("参数解析失败：%s", err)
		common.FailWithError(err, &cr, c)
		return
	}

	//获取标签的 id
	id := c.Param("id")

	var tag models.TagModel
	err = global.DB.Take(&tag, id).Error
	if err != nil {
		common.FailWithMessage("标签不存在！", c)
		return
	}

	//global.Log.Warnln(cr)

	/*
		可以使用一个struct转map的第三方库：structs,很方便的实现struct对map的转化
		使用 go get github.com/fatih/structs
	*/

	//修改数据
	err = global.DB.Model(&models.TagModel{}).Where("id = ?", id).Updates(structs.Map(&cr)).Error

	if err != nil {
		global.Log.Error("更新标签失败：", err)
		common.FailWithMessage("更新标签失败！", c)
		return
	}

	common.OKWithMessage("更新标签成功！", c)

}
