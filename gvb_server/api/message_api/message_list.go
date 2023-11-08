package message_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/utils/jwts"
)

func (MessageApi) MessageListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var messageList []models.MessageModel
	global.DB.Find(&messageList, "send_user_id = ? or rev_user_id = ?", claims.ID, claims.ID)

	global.Log.Warnln(claims.ID)
	global.Log.Warnln(messageList)

	common.OKWithData(messageList, c)
}
