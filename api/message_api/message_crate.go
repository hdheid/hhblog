package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
)

type MessageRequest struct {
	SendUserID uint   `json:"send_user_id" binding:"required"` //发送人id
	RevUserID  uint   `json:"rev_user_id" binding:"required"`  //接收人id
	Content    string `json:"content" binging:"required"`      //消息内容
}

// MessageCreateView 发布消息
func (MessageApi) MessageCreateView(c *gin.Context) {
	var cr MessageRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Debugf("参数解析失败:%s", err)
		common.FailWithError(err, &cr, c)
		return
	}

	//SendUserID 就是当前登录人的ID

	//查询发送人和接收人
	var sendUser models.UserModel
	err = global.DB.Take(&sendUser, cr.SendUserID).Error
	if err != nil {
		global.Log.Errorf("查询发送人失败 %s", err)
		common.FailWithMessage("查询发送人失败", c)
		return
	}

	var revUser models.UserModel
	err = global.DB.Take(&revUser, cr.RevUserID).Error
	if err != nil {
		global.Log.Errorf("查询接收人失败 %s", err)
		common.FailWithMessage("查询接收人失败", c)
		return
	}

	//将消息信息入库
	err = global.DB.Create(&models.MessageModel{
		SendUserID:       cr.SendUserID,
		SendUserNickName: sendUser.NickName,
		SendUserAvatar:   sendUser.Avatar,
		RevUserID:        cr.RevUserID,
		RevUserNickName:  revUser.NickName,
		RevUserAvatar:    revUser.Avatar,
		IsRead:           false,
		Content:          cr.Content,
	}).Error
	if err != nil {
		global.Log.Errorf("消息入库失败 %s", err)
		common.FailWithMessage("消息发送失败", c)
		return
	}

	common.OKWithMessage("消息发送成功！", c)
}
