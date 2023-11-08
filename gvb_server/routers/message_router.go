package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func MessageRouter(r *gin.RouterGroup) {
	MessageApi := api.ApiGroupApp.MessageApi
	r.POST("/messages", MessageApi.MessageCreateView)
}
