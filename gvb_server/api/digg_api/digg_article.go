package digg_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/service/redis_ser"
)

//点赞逻辑有问题，并没有将点赞数据同步到es，而且点赞数还在文章列表里面进行操作，逻辑有点混乱，后期优化

func (DiggApi) DiggArticleView(c *gin.Context) {
	var cr models.ESIdRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err.Error())
		common.FailWithCode(common.ArgumentError, c)
		return
	}

	err = redis_ser.NewDigg().Set(cr.ID)
	if err != nil {
		global.Log.Error(err.Error())
		common.FailWithMessage("点赞失败！", c)
	}

	common.OKWithMessage("点赞成功！", c)
}
