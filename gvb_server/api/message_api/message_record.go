package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/utils/jwts"
)

type MessageRecordRequest struct {
	UserID uint `json:"user_id" binging:"required" msg:"请输入查询的用户id"`
}

func (MessageApi) MessageRecordView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var cr MessageRecordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Debugf("参数解析失败：%s", err)
		common.FailWithCode(common.ArgumentError, c)
		return
	}

	var messageList []models.MessageModel
	//按照从早到晚的顺序将数据全部取出来
	//发送人和接收人的id和如果相同，就表示这是一组消息（正确性待验证，后续可以看看其他项目怎么的）

	global.DB.Order("created_at asc").
		Where("send_user_id = ? and rev_user_id = ?", claims.UserID, cr.UserID).
		Or("send_user_id = ? and rev_user_id = ?", cr.UserID, claims.UserID).
		Find(&messageList)

	//原代码
	//var _messageList []models.MessageModel
	//var messageList = make([]models.MessageModel, 0)
	//global.DB.Order("created_at asc").
	//	Find(&_messageList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID)
	//for _, model := range _messageList {
	//	if model.RevUserID == cr.UserID || model.SendUserID == cr.UserID {
	//		messageList = append(messageList, model)
	//	}
	//}

	if messageList == nil {
		messageList = make([]models.MessageModel, 0)
	}

	common.OKWithData(messageList, c)
}
