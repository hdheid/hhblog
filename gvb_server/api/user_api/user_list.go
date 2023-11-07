package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common"
	"gvb_server/models/ctype"
	"gvb_server/service/list_func"
	"gvb_server/utils/desens"
	"gvb_server/utils/jwts"
)

func (UserApi) UserListView(c *gin.Context) {
	//如何判断用户的权限？使用token，通过中间件地上下文信息获取token
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims) //断言

	var conf models.PageInfo
	err := c.ShouldBindQuery(&conf)
	if err != nil {
		global.Log.Debugf("参数解析失败：%s", err)
		common.FailWithCode(common.ArgumentError, c)
		return
	}

	userList, count, _ := list_func.ComList(models.UserModel{}, list_func.Option{
		PageInfo: conf,
	})

	//对用户的数据进行一定地脱敏处理
	var users []models.UserModel
	for _, user := range userList {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin { //如果当前登录用户不是管理员，则不对其展示所有人的用户名，只展示昵称
			user.UserName = ""
		}

		user.Email = desens.DesensitizationEmail(user.Email)
		user.Tel = desens.DesensitizationTel(user.Tel)

		users = append(users, user)
	}

	common.OKWithList(users, count, c) //返回响应
}
