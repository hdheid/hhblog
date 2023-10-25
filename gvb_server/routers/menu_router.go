package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func MenuRouter(r *gin.RouterGroup) {
	menuApi := api.ApiGroupApp.MenuApi
	r.POST("/menus", menuApi.MenuCreateView)
	r.GET("/menus", menuApi.MenuListView)
	r.GET("/menus_names", menuApi.MenuNameListView)
	r.GET("/menus/:id", menuApi.MenuDetailView)
	r.PUT("/menus/:id", menuApi.MenuUpdateView)
	r.DELETE("/menus", menuApi.MenuDeleteView)
}
