package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func SettingsRouter(r *gin.RouterGroup) {
	settingsApi := api.ApiGroupApp.SettingsApi
	r.GET("/settings", settingsApi.SettingsInfoView)
}
