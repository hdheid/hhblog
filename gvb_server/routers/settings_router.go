package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func SettingsRouter(r *gin.RouterGroup) {
	settingsApi := api.ApiGroupApp.SettingsApi
	r.GET("/settings/site", settingsApi.SettingsSiteInfoView)
	r.PUT("/settings/site", settingsApi.SettingsSiteUpdateView)
	r.GET("/settings/:name", settingsApi.SettingsInfoView)       //四个信息
	r.PUT("/settings/:name", settingsApi.SettingsInfoUpdateView) //修改四个信息
}
