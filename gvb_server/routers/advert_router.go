package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func AdvertRouter(r *gin.RouterGroup) {
	advertApi := api.ApiGroupApp.AdvertApi
	r.POST("/adverts", advertApi.AdvertCreateView)
	r.GET("/adverts", advertApi.AdvertListView)
	r.PUT("/adverts/:id", advertApi.AdvertUpdateView)
	r.DELETE("/adverts", advertApi.AdvertRemoveView)
}
