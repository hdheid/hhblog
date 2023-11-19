package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

func MessageRouter(r *gin.RouterGroup) {
	MessageApi := api.ApiGroupApp.MessageApi
	r.POST("/messages", middleware.JwtAuth(), MessageApi.MessageCreateView)
	r.GET("/messages_all", middleware.JwtAuth(), MessageApi.MessageListAllView)
	r.GET("/messages", middleware.JwtAuth(), MessageApi.MessageListView)
	r.GET("/messages_record", middleware.JwtAuth(), MessageApi.MessageRecordView)

}
