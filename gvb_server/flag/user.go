package flag

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/utils"
)

func CreateUser(permissions string) {
	// 创建用户的逻辑
	// 用户名 昵称 密码 确认密码 邮箱
	var (
		userName   string
		nickName   string
		password   string
		rePassword string
		email      string
	)
	fmt.Printf("请输入用户名：")
	fmt.Scan(&userName)
	fmt.Printf("请输入昵称：")
	fmt.Scan(&nickName)
	fmt.Printf("请输入邮箱：")
	fmt.Scan(&email)
	fmt.Printf("请输入密码：")
	fmt.Scan(&password)
	fmt.Printf("请再次输入密码：")
	fmt.Scan(&rePassword)

	//判断输入的用户名是否存在
	var user models.UserModel
	err := global.DB.Take(&user, "user_name = ?", userName).Error
	if err == nil {
		//存在
		global.Log.Error("用户名已存在，请从新输入！")
		return
	}

	//校验密码
	if password != rePassword {
		global.Log.Error("两次密码不一致，请从新输入！")
		return
	}

	//对密码哈希传入数据库
	hashPwd := utils.HashPwd(password)

	//头像：1. 默认头像 2. 随机头像
	avatar := "/static/avatar/default.jpg" //默认头像

	//入库
	role := ctype.PermissionUser
	if permissions == "admin" {
		role = ctype.PermissionAdmin
	}

	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		Avatar:     avatar,
		IP:         "127.0.0.1",
		Addr:       "本地地址",
		SignStatus: ctype.SignEmail,
	}).Error

	if err != nil {
		global.Log.Error("用户创建失败！", err)
		return
	}
	global.Log.Infof("用户 %s 创建成功!", userName)
}
