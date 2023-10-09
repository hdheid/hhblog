package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common"
)

/*
ShouldBind：它能够基于请求自动提取JSON、form表单和QueryString型的数据，并把值绑定到指定的结构体对象

显示某一项的配置信息
*/

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var query SettingsUri
	err := c.ShouldBindUri(&query) //从前端传递过来动态路由的name绑定到query上面
	if err != nil {
		global.Log.Debug("绑定数据到query上失败，", err)
		common.FailWithCode(common.ArgumentError, c) //参数错误
		return
	}

	switch query.Name {
	case "email":
		common.OKWithData(global.Config.Email, c)
	case "qq":
		common.OKWithData(global.Config.QQ, c)
	case "qiniu":
		common.OKWithData(global.Config.QiNiu, c)
	case "jwt":
		common.OKWithData(global.Config.Jwy, c)
	default:
		common.FailWithMessage("没有对应配置！", c)
	}
}
