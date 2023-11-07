package user_api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/plugins/email"
	"gvb_server/utils"
	"gvb_server/utils/jwts"
	"gvb_server/utils/random"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required" msg:"邮箱非法"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

// UserBindEmailView 用户绑定邮箱
func (UserApi) UserBindEmailView(c *gin.Context) {
	/*
		1. 首先用户会输入邮箱
		2. 后台会给用户发送验证码
		3. 用户输入验证码，密码
		4. 验证成功，完成绑定
	*/
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var cr BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		common.FailWithError(err, &cr, c)
		return
	}

	session := sessions.Default(c)
	if cr.Code == nil {
		//第一次发验证码
		//生成4位验证码，将验证码存入session
		code := random.GenValidateCode(4)
		//写入session
		/*
			在这里session的作用就是将第一次调用该接口的时候生成的code存下来，用于再次调用该接口的时候进行对比验证。
			因为如果不存起来，那么第一次调用的时候生成的code就会消失，因此，如果使用redis来存储也是可以的。
		*/
		session.Set("valid_code", code)
		err = session.Save()
		if err != nil {
			global.Log.Error(err)
			common.FailWithMessage("session错误！", c)
			return
		}

		err = email.NewCode().Send(cr.Email, "你的验证码是："+code)
		if err != nil {
			global.Log.Error(err)
		}

		common.OKWithMessage("验证码已发送！", c)
		return
	}

	code := session.Get("valid_code")
	if code != *cr.Code {
		common.FailWithMessage("验证码错误！", c)
		return
	}

	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		global.Log.Error(err)
		common.FailWithMessage("该用户不存在！", c)
		return
	}

	/*
		这里至少会调用两次接口，第一次用于生成验证码，第二次用于校验验证码。
		因此，两次传过来的邮箱地址需要进行一致性校验，否则可能会出问题
	*/

	//在前端可以对输入的密码进行长度和强度的限制，如果前端没有这个功能，那么就在这里写
	hashPwd := utils.HashPwd(cr.Password)
	err = global.DB.Model(&user).Updates(map[string]any{
		"email":    cr.Email,
		"password": hashPwd,
	}).Error
	if err != nil {
		global.Log.Error(err)
		common.FailWithMessage("绑定邮箱失败！", c)
		return
	}

	common.OKWithMessage("邮箱绑定成功！", c)
}
