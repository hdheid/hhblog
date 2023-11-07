package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/utils"
	"gvb_server/utils/jwts"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入用密码"`
}

// EmailLoginView 邮箱用户名登录
func (UserApi) EmailLoginView(c *gin.Context) {
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	//global.Log.Warnln(err)
	if err != nil {
		global.Log.Debugf("参数解析失败:%s", err)
		common.FailWithError(err, &cr, c)
		return
	}

	//验证用户是否存在
	var user models.UserModel
	err = global.DB.Take(&user, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		global.Log.Warnln("用户名不存在！")
		common.FailWithMessage("用户名或密码错误！", c)
		return
	}

	//判断密码是否正确
	isCheck := utils.CheckPwd(user.Password, cr.Password)
	if !isCheck {
		global.Log.Warnln("密码验证错误！")
		common.FailWithMessage("用户名或密码错误！", c)
		return
	}

	//登录成功的话，生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: user.NickName,
		Role:     int(user.Role),
		UserID:   user.ID,
		Avatar:   user.Avatar,
	})
	if err != nil {
		global.Log.Error("生成token失败！", err)
		common.FailWithMessage("生成token失败！", c)
		return
	}

	common.OKWithData(token, c)
}
