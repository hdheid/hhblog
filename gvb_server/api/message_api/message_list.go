package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/utils/jwts"
	"time"
)

type Message struct {
	SendUserID       uint      `json:"send_user_id"` // 发送人id
	SendUserNickName string    `json:"send_user_nick_name"`
	SendUserAvatar   string    `json:"send_user_avatar"`
	RevUserID        uint      `json:"rev_user_id"` // 接收人id
	RevUserNickName  string    `json:"rev_user_nick_name"`
	RevUserAvatar    string    `json:"rev_user_avatar"`
	Content          string    `json:"content"`       // 消息内容
	CreatedAt        time.Time `json:"created_at"`    // 最新的消息时间
	MessageCount     int       `json:"message_count"` // 消息条数
}

type MessageGroup map[uint]*Message

func (MessageApi) MessageListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var messageGroup = MessageGroup{}
	var messageList []models.MessageModel
	var messages = make([]Message, 0)

	global.DB.Order("created_at asc").
		Find(&messageList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID)
	//按照从早到晚的顺序将数据全部取出来
	//发送人和接收人的id和如果相同，就表示这是一组消息（正确性待验证，后续可以看看其他项目怎么的）

	for _, model := range messageList {
		message := Message{
			SendUserID:       model.SendUserID,
			SendUserNickName: model.SendUserNickName,
			SendUserAvatar:   model.SendUserAvatar,
			RevUserID:        model.RevUserID,
			RevUserNickName:  model.RevUserNickName,
			RevUserAvatar:    model.RevUserAvatar,
			Content:          model.Content,
			CreatedAt:        model.CreatedAt,
			MessageCount:     1,
		}

		idNum := model.SendUserID + model.RevUserID
		val, ok := messageGroup[idNum]
		if !ok {
			//如果不存在这组消息，就更新一下
			messageGroup[idNum] = &message
			continue
		}
		//存在则将其更新为最新的时间的消息
		message.MessageCount = val.MessageCount + 1
		messageGroup[idNum] = &message
	}

	for _, message := range messageGroup {
		messages = append(messages, *message)
	}

	common.OKWithData(messages, c)
}
