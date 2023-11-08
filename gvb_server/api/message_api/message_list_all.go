package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/service/list_func"
)

func (MessageApi) MessageListAllView(c *gin.Context) {
	var conf models.PageInfo
	err := c.ShouldBindQuery(&conf)
	if err != nil {
		global.Log.Debugf("参数解析失败：%s", err)
		common.FailWithCode(common.ArgumentError, c)
		return
	}

	tagList, count, _ := list_func.ComList(models.MessageModel{}, list_func.Option{
		PageInfo: conf,
	})

	common.OKWithList(tagList, count, c) //返回响应
}
