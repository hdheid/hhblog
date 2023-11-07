package user_ser

import (
	"errors"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/utils"
)

func (UserService) CreateUser(userName, nickName, password, ip string, role ctype.Role, email string) error {
	//判断输入的用户名是否存在
	var user models.UserModel
	err := global.DB.Take(&user, "user_name = ?", userName).Error
	if err == nil {
		//存在
		global.Log.Error("用户名已存在，请从新输入！")
		return errors.New("用户名已存在，请从新输入！")
	}

	//对密码哈希传入数据库
	hashPwd := utils.HashPwd(password)

	//头像：1. 默认头像 2. 随机头像
	avatar := "/static/avatar/default.jpg" //默认头像

	//通过ip获取地址
	addr := GetLocation(ip)
	if ip == "127.0.0.1" {
		addr = "本地地址"
	}

	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		Avatar:     avatar,
		IP:         ip,
		Addr:       addr,
		SignStatus: ctype.SignEmail,
	}).Error

	if err != nil {
		return err
	}

	return nil
}
