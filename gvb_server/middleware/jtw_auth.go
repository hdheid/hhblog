package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common"
	"gvb_server/utils/jwts"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			common.FailWithMessage("未携带token！", c)
			c.Abort()
			return
		}

		claims, err := jwts.ParseToken(token)
		if err != nil {
			common.FailWithMessage("token错误！", c)
			c.Abort()
			return
		}

		//拿到token后，看看该用户是不是已经注销了，如果注销了，就不能进行登录用户才能干的事情
		//直接判断该token是不是在Redis中存在即可
		// 使用Redis的Exists命令判断键是否存在
		exists, err := global.RDB.Exists(fmt.Sprintf("logout_%s", token)).Result()
		global.Log.Info("相较于原代码优化的点")
		if err != nil {
			// 处理错误
			common.FailWithMessage("查找token失败！", c)
			c.Abort()
			return
		}
		if exists == 1 { //如果在 redis 中找到了，就表示用户已经注销过了，已经失效了
			common.FailWithMessage("token 已经失效！", c)
			c.Abort()
			return
		}

		//登录的用户信息，设置上下文信息
		c.Set("claims", claims)
	}
}
