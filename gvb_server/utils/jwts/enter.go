package jwts

import "github.com/dgrijalva/jwt-go/v4"

type JwtPayLoad struct {
	NickName string `json:"nick_name"` //用户昵称
	Role     int    `json:"role"`      //权限 1:普通用户 2:管理员 3:游客
	UserID   uint   `json:"user_id"`   //用户id
	Avatar   string `json:"avatar"`    //用户头像
}

var MySecret []byte

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}
