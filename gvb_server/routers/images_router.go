package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func ImagesRouter(r *gin.RouterGroup) {
	imagesApi := api.ApiGroupApp.ImagesApi
	r.POST("/images", imagesApi.ImageUploadView)
	r.GET("/images", imagesApi.ImageListView)
	r.GET("/images_names", imagesApi.ImageNameListView)
	r.DELETE("/images", imagesApi.ImageRemoveView)
	r.PUT("/images", imagesApi.ImageUpdateView)
}
