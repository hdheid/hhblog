package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

func DiggRouter(r *gin.RouterGroup) {
	diggApi := api.ApiGroupApp.Digg
	r.POST("/digg/article", middleware.JwtAdmin(), diggApi.DiggArticleView)
}
