package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func TagRouter(r *gin.RouterGroup) {
	TagApi := api.ApiGroupApp.TagApi
	r.POST("/tags", TagApi.TagCreateView)
	r.GET("/tags", TagApi.TagListView)
	r.PUT("/tags/:id", TagApi.TagUpdateView)
	r.DELETE("/tags", TagApi.TagRemoveView)
}
