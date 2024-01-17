package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

func DiggRouter(r *gin.RouterGroup) {
	diggApi := api.ApiGroupApp.DiggApi
	r.POST("/digg/article", middleware.JwtAdmin(), diggApi.DiggArticleView)
}
