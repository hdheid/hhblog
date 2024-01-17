package tag_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
)

func (TagApi) TagRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Debugf("参数解析失败：%s", err)
		common.FailWithCode(common.ArgumentError, c)
		return
	}

	var tagList []models.TagModel
	count := global.DB.Find(&tagList, cr.IDList).RowsAffected
	if count == 0 {
		common.FailWithMessage("标签不存在！", c)
		return
	}

	global.DB.Delete(&tagList) //删除数据库中的相应广告信息
	common.OKWithMessage(fmt.Sprintf("一共删除了%d个标签！", count), c)
}
