package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common"
	"gvb_server/utils/jwts"
	"time"
)

func (UserApi) UserLogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims) //断言
	token := c.Request.Header.Get("token")

	//这里的 ExpiresAt 是截至时间，而 Redis 的时间是多少秒之后过期，如果需要存储带Redis中，需要进行转换
	exp := claims.ExpiresAt
	now := time.Now()

	diff := exp.Time.Sub(now)
	err := global.RDB.Set(fmt.Sprintf("logout_%s", token), "ExpirationTime", diff).Err() //储存在Redis中
	if err != nil {
		global.Log.Errorf("Redis 连接失败 %s", err)
		common.FailWithMessage("注销失败！", c)
		return
	}

	common.OKWithMessage("注销成功！", c)
}
