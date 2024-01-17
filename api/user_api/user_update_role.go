package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/models/ctype"
)

type UserRoleRequest struct {
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误！"`
	NickName string     `json:"nick_name"` //管理员可以修改用户昵称，防止用户昵称非法
	UserID   uint       `json:"user_id" binding:"required" msg:"用户id错误"`
}

// UserUpdateRoleView 用户权限变更
func (UserApi) UserUpdateRoleView(c *gin.Context) {
	var cr UserRoleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		common.FailWithError(err, &cr, c)
		return
	}

	//修改用户权限
	var user models.UserModel
	err = global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		global.Log.Debug("查找用户ID失败，用户不存在！")
		common.FailWithMessage("用户id错误，用户不存在！", c)
		return
	}

	err = global.DB.Model(&user).Updates(map[string]any{
		"role":      cr.Role,
		"nick_name": cr.NickName,
	}).Error
	if err != nil {
		global.Log.Debug("用户信息变更失败！")
		common.FailWithMessage("用户信息变更失败！", c)
		return
	}

	common.OKWithMessage("用户信息变更成功！", c)
}
