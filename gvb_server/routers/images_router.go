package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func ImagesRouter(r *gin.RouterGroup) {
	imagesApi := api.ApiGroupApp.ImagesApi
	r.POST("/images", imagesApi.ImageUploadView)
}
