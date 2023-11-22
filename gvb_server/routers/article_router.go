package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

func ArticleRouter(r *gin.RouterGroup) {
	articleApi := api.ApiGroupApp.ArticleApi
	r.POST("/articles", middleware.JwtAdmin(), articleApi.ArticleCreateView)
}
