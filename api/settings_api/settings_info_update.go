package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/common"
)

/*
修改某一项的配置信息
*/

func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var query SettingsUri
	err := c.ShouldBindUri(&query) //从前端传递过来动态路由的name绑定到query上面
	if err != nil {
		global.Log.Debug("绑定数据到query上失败，", err)
		common.FailWithCode(common.ArgumentError, c) //参数错误
		return
	}

	switch query.Name {
	case "email":
		var email config.Email
		err := c.ShouldBindJSON(&email) //从前端传递过来的修改后的数据，绑定到conf上
		if err != nil {
			global.Log.Debug("绑定数据到conf上失败，", err)
			common.FailWithCode(common.ArgumentError, c) //参数错误
			return
		}
		global.Config.Email = email //由于Config是指针形式传递的，所以这样可以直接修改值

	case "qq":
		var qq config.QQ
		err := c.ShouldBindJSON(&qq) //从前端传递过来的修改后的数据，绑定到conf上
		if err != nil {
			global.Log.Debug("绑定数据到conf上失败，", err)
			common.FailWithCode(common.ArgumentError, c) //参数错误
			return
		}
		global.Config.QQ = qq //由于Config是指针形式传递的，所以这样可以直接修改值

	case "qiniu":
		var qiniu config.QiNiu
		err := c.ShouldBindJSON(&qiniu) //从前端传递过来的修改后的数据，绑定到conf上
		if err != nil {
			global.Log.Debug("绑定数据到conf上失败，", err)
			common.FailWithCode(common.ArgumentError, c) //参数错误
			return
		}
		global.Config.QiNiu = qiniu //由于Config是指针形式传递的，所以这样可以直接修改值

	case "jwt":
		var jwt config.Jwy
		err := c.ShouldBindJSON(&jwt) //从前端传递过来的修改后的数据，绑定到conf上
		if err != nil {
			global.Log.Debug("绑定数据到conf上失败，", err)
			common.FailWithCode(common.ArgumentError, c) //参数错误
			return
		}
		global.Config.Jwy = jwt //由于Config是指针形式传递的，所以这样可以直接修改值

	default:
		common.FailWithMessage("没有对应配置！", c)
	}

	err = core.SetYaml() //修改配置文件
	if err != nil {
		global.Log.Error("修改配置文件失败，", err)
		common.FailWithMessage(err.Error(), c) //将失败信息响应回去
		return
	}
	common.OKWithNoting(c)
}
