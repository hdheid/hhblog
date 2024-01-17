package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/utils"
	"gvb_server/utils/jwts"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd" binding:"required" msg:"请输入原密码"`
	NewPwd string `json:"new_pwd" binding:"required" msg:"请输入新密码"`
}

// UserUpdatePassword 修改密码
func (UserApi) UserUpdatePassword(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var cr UpdatePasswordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		common.FailWithError(err, &cr, c)
		return
	}

	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		global.Log.Debug("查找用户ID失败，用户不存在！")
		common.FailWithMessage("该用户不存在!", c)
		return
	}

	//修改密码环节
	//首先判断密码是否一致
	if !utils.CheckPwd(user.Password, cr.OldPwd) {
		global.Log.Debug("密码错误!")
		common.FailWithMessage("密码错误!", c)
		return
	}

	//更新密码
	newHashPwd := utils.HashPwd(cr.NewPwd)
	err = global.DB.Model(&user).Update("password", newHashPwd).Error
	if err != nil {
		global.Log.Error("密码修改失败，请联系管理员！", err)
		common.FailWithMessage("密码修改失败，请联系管理员！", c)
		return
	}

	common.OKWithMessage("修改密码成功！", c)
}
