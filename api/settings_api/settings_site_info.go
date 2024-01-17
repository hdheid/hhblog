package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common"
)

func (SettingsApi) SettingsSiteInfoView(c *gin.Context) {
	common.OKWithData(global.Config.SiteInfo, c)
}
