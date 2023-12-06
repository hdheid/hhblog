package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

func CommentRouter(r *gin.RouterGroup) {
	commentApi := api.ApiGroupApp.CommentApi
	r.POST("/comments", middleware.JwtAdmin(), commentApi.CommentCreateView)
	r.GET("/comments", commentApi.CommentList)
}
