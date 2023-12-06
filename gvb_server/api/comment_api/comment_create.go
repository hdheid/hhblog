package comment_api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/service/es_ser"
	"gvb_server/service/redis_ser"
	"gvb_server/utils/jwts"
)

type CommentRequest struct {
	ArticleID       string `json:"article_id" binding:"required" msg:"请选择文章"`
	Content         string `json:"content" binding:"required" msg:"请输入评论内容"`
	ParentCommentID *uint  `json:"parent_comment_id"` //父评论 id
}

func (CommentApi) CommentCreateView(c *gin.Context) {
	var cr CommentRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err.Error())
		common.FailWithError(err, &cr, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	//文章是否存在
	_, err = es_ser.CommonDetail(cr.ArticleID)
	if err != nil {
		global.Log.Error("文章不存在！", claims.UserID)
		common.FailWithMessage("文章不存在！", c)
		return
	}

	//判断是否是子评论
	if cr.ParentCommentID != nil {
		//表示是子评论
		//给父评论数加一
		var parentComment models.CommentModel
		err = global.DB.Take(&parentComment, cr.ParentCommentID).Error //首先找一下是否有这个父评论
		if err != nil {
			common.FailWithMessage("父评论不存在！", c)
			return
		}

		//判断找到的父评论是不是这篇文章的评论
		if parentComment.ArticleID != cr.ArticleID {
			common.FailWithMessage("评论文章不一致！", c)
			return
		}

		//给父评论的子评论个数加一
		global.DB.Model(&parentComment).Update("comment_count", gorm.Expr("comment_count + ?", 1))
	}

	//添加评论
	global.DB.Create(&models.CommentModel{
		ParentCommentID: cr.ParentCommentID,
		Content:         cr.Content,
		ArticleID:       cr.ArticleID,
		UserID:          claims.UserID,
	})

	//拿到文章数，新的文章评论数存在缓存里面
	redis_ser.NewCommentCount().Set(cr.ArticleID)

	common.OKWithMessage("评论成功！", c)
}
