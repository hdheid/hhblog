package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/common"
)

func (SettingsApi) SettingsSiteUpdateView(c *gin.Context) {
	var conf config.SiteInfo
	err := c.ShouldBindJSON(&conf) //从前端传递过来的修改后的数据，绑定到conf上
	if err != nil {
		global.Log.Debug("绑定数据到conf上失败，", err)
		common.FailWithCode(common.ArgumentError, c) //参数错误
		return
	}

	global.Config.SiteInfo = conf //由于Config是指针形式传递的，所以这样可以直接修改值
	err = core.SetYaml()          //修改配置文件
	if err != nil {
		global.Log.Error("修改配置文件失败，", err)
		common.FailWithMessage(err.Error(), c) //将失败信息响应回去
		return
	}
	common.OKWithNoting(c)
}
