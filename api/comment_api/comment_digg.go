package comment_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/service/redis_ser"
)

type CommentIDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

func (CommentApi) CommentDiggView(c *gin.Context) {
	var cr CommentIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		global.Log.Error(err.Error())
		common.FailWithCode(common.ArgumentError, c)
		return
	}

	//查一下文章是否存在
	var comment models.CommentModel
	err = global.DB.Take(&comment, cr.ID).Error
	if err != nil {
		common.FailWithMessage("评论不存在！", c)
		return
	}

	//点赞操作
	redis_ser.NewCommentDigg().Set(fmt.Sprintf("%d", cr.ID))
	common.OKWithMessage("点赞成功！", c)
}
