package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common"
	"gvb_server/models/ctype"
	"gvb_server/service/user_ser"
)

type UserCreateRequest struct {
	NickName string     `json:"nick_name" binding:"required" msg:"请输入昵称"`  //昵称
	UserName string     `json:"user_name" binding:"required" msg:"请输入用户名"` //用户名
	Password string     `json:"password" binding:"required" msg:"请输入密码"`   //密码
	Role     ctype.Role `json:"role" binding:"required" msg:"请选择权限"`       //权限 1 管理员  2 普通用户  3 游客
}

func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Debugf("参数解析失败:%s", err)
		common.FailWithError(err, &cr, c)
		return
	}

	err = user_ser.UserService{}.CreateUser(cr.UserName, cr.NickName, cr.Password, c.ClientIP(), cr.Role, "")
	//global.Log.Warnln(c.ClientIP())

	if err != nil {
		global.Log.Error("用户创建失败！", err)
		common.FailWithMessage(err.Error(), c)
		return
	}
	global.Log.Infof("用户 %s 创建成功!", cr.UserName)
	common.OKWithMessage(fmt.Sprintf("用户%s创建成功!", cr.UserName), c)
}
