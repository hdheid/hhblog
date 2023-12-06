//package main
//
//import (
//	"fmt"
//	"gorm.io/gorm"
//	"gvb_server/core"
//	"gvb_server/global"
//	"gvb_server/models"
//)
//
//func main() {
//	core.InitConf()
//	core.InitES()
//	core.InitGorm()
//	FindArticleCommentList("4Hk-JYwBphBkeiGY9bi_")
//}
//
//func FindArticleCommentList(articleID string) {
//	//先查出根评论，也就是所有的父评论
//	var RootCommentList []*models.CommentModel
//	global.DB.Find(&RootCommentList, "article_id = ? and parent_comment_id is null", articleID)
//
//	//遍历根评论，递归查询所有子评论
//	for _, model := range RootCommentList {
//		var subCommentList []models.CommentModel
//		FindSubComment(*model, &subCommentList)
//		fmt.Println(subCommentList)
//	}
//}
//
//// FindSubComment 递归查询子评论
//func FindSubComment(model models.CommentModel, subCommentList *[]models.CommentModel) {
//	err := global.DB.Preload("SubComments").Take(&model).Error
//	if err == gorm.ErrRecordNotFound {
//		//没找到
//		return
//	}
//	for _, sub := range model.SubComments {
//		*subCommentList = append(*subCommentList, sub)
//		FindSubComment(sub, subCommentList)
//	}
//	return
//}

package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
)

func main() {
	core.InitConf()
	//core.InitES()
	core.InitGorm()
	FindArticleCommentList("4Hk-JYwBphBkeiGY9bi_")
}

func FindArticleCommentList(articleID string) {
	// 先把文章下的根评论查出来
	var RootCommentList []*models.CommentModel
	global.DB.Find(RootCommentList, "article_id = ? and parent_comment_id is null", articleID)
	// 遍历根评论，递归查根评论下的所有子评论
	//for _, model := range RootCommentList {
	//fmt.Println(model)
	//var subCommentList []models.CommentModel
	//FindSubComment(*model, &subCommentList)
	//model.SubComments = subCommentList
	//}
}

// FindSubComment 递归查评论下的子评论
//func FindSubComment(model models.CommentModel, subCommentList *[]models.CommentModel) {
//	global.DB.Preload("SubComments").Take(&model)
//	for _, sub := range model.SubComments {
//		*subCommentList = append(*subCommentList, sub)
//		FindSubComment(sub, subCommentList)
//	}
//	return
//}
