package utils

import (
	"golang.org/x/crypto/bcrypt"
	"gvb_server/global"
)

// HashPwd 密码加密
func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		global.Log.Warnln("密码加密失败！", err)
	}
	return string(hash)
}

func CheckPwd(hashPwd string, pwd string) bool {
	byteHash := []byte(hashPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, []byte(pwd))
	if err != nil {
		global.Log.Warnln("密码验证失败！", err)
		return false
	}

	return true
}
