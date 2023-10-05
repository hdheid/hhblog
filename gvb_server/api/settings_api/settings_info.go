package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/common"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	common.FailWitCode(common.SettingsErr, c)
}
